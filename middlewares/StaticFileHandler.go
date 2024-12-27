package middlewares

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
)

const INDEX = "index.html"

type ServeFileSystem interface {
	http.FileSystem
	Exists(prefix string, path string) bool
}

type localFileSystem struct {
	http.FileSystem
	root    string
	indexes bool
}

func LocalFile(root string, indexes bool) *localFileSystem {
	return &localFileSystem{
		FileSystem: gin.Dir(root, indexes),
		root:       root,
		indexes:    indexes,
	}
}

func (l *localFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		name := path.Join(l.root, p)
		stats, err := os.Stat(name)
		if err != nil {
			return false
		}
		if stats.IsDir() {
			if !l.indexes {
				index := path.Join(name, INDEX)
				_, err := os.Stat(index)
				if err != nil {
					return false
				}
			}
		}
		return true
	}
	return false
}

func ServeRoot(urlPrefix, root string) gin.HandlerFunc {
	return Serve(urlPrefix, LocalFile(root, false))
}

// Static returns a middleware handler that serves static files in the given directory.
func Serve(urlPrefix string, fs ServeFileSystem) gin.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		if fs.Exists(urlPrefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

// EmbedFolder 从一个嵌入的文件系统中提取指定的文件夹，并返回一个 ServeFileSystem 接口实例。
// 这个函数的目的是为了在嵌入式文件系统中定位并提供指定路径下的文件和目录，以便通过 HTTP 服务进行提供。
// 参数:
//
//	fsEmbed - 嵌入式文件系统的嵌入对象，通常通过 go:embed 导入。
//	targetPath - 目标文件夹的路径，该路径是相对于嵌入式文件系统的根目录的。
//
// 返回值:
//
//	ServeFileSystem - 一个接口类型，用于提供文件服务。这个接口通常用于 http.FileServer 函数，以提供 HTTP 文件服务。
func EmbedFolder(fsEmbed embed.FS, targetPath string) ServeFileSystem {
	// 从嵌入式文件系统中提取指定路径的文件系统。
	fsys, err := fs.Sub(fsEmbed, targetPath)
	// 如果提取过程中发生错误，则抛出 panic。
	if err != nil {
		panic(err)
	}
	// 返回一个 embedFileSystem 结构体实例，它实现了 ServeFileSystem 接口。
	// 这个实例包装了提取出来的文件系统，使其可以通过 HTTP 服务提供文件。
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
