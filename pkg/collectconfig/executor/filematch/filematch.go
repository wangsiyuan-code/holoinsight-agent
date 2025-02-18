/*
 * Copyright 2022 Holoinsight Project Authors. Licensed under Apache-2.0.
 */

package filematch

import "github.com/traas-stack/holoinsight-agent/pkg/collectconfig/executor/logstream"

const (
	TypePath   = "path"
	TypeGlob   = "glob"
	TypeRegexp = "regexp"
	TypeFormat = "format"
	TypeSls    = "sls"
)

type (
	FileMatcher interface {
		Find() ([]FatPath, int, error)
		// 是否动态匹配多文件?
		IsDynamicMultiFiles() bool
	}
	// 富路径
	// 对于绝对路径来说, tags==nil
	// 对于 /home/admin/{a}/{b} 的路径来说 会有2个tags
	FatPath struct {
		Path      string
		Tags      map[string]string
		IsSls     bool
		SlsConfig logstream.SlsConfig
	}
)

func GetPaths(fatPaths []FatPath) []string {
	ret := make([]string, len(fatPaths))
	for i := range fatPaths {
		ret[i] = fatPaths[i].Path
	}
	return ret
}
