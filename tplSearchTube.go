// Copyright 2016 - 2020 The aurora Authors. All rights reserved. Use of this
// source code is governed by a MIT license that can be found in the LICENSE
// file.
//
// The aurora is a web-based beanstalkd queue server console written in Go
// and works on macOS, Linux and Windows machines. Main idea behind using Go
// for backend development is to utilize ability of the compiler to produce
// zero-dependency binaries for multiple platforms. aurora was created as an
// attempt to build very simple and portable application to work with local or
// remote beanstalkd server.

package main

import (
	"strconv"
	"strings"
)

// tplSearchTube rander navigation search box for search content in jobs by given tube.
func tplSearchTube(server string, tube string, state string) string {
	buf := strings.Builder{}
	buf.WriteString(`<form class="navbar-form navbar-right" style="margin-top:5px;margin-bottom:0px;" role="search" method="get"><input type="hidden" name="server" value="`)
	buf.WriteString(server)
	buf.WriteString(`"/><input type="hidden" name="tube" value="`)
	buf.WriteString(tube)
	buf.WriteString(`"/><input type="hidden" name="action" value="search"/><input type="hidden" name="limit" value="`)
	buf.WriteString(strconv.Itoa(selfConf.SearchResultLimit))
	buf.WriteString(`"/><div class="form-group"><select class="form-control input-sm search-query" style="margin-right: 10px;" id="state" name="state"><option value="" disabled selected>Select state</option><option>ready</option><option>delayed</option><option>buried</option></select>`)
	buf.WriteString(`<input type="text" class="form-control input-sm search-query" style="margin-right: 10px;" name="id" placeholder="Bean ID"/>`)
	buf.WriteString(`<input type="text" class="form-control input-sm search-query" style="margin-right: 10px;" name="searchStr" placeholder="Search this tube"/>`)
	buf.WriteString(`<button type="submit" class="btn btn-info btn-sm">Search</button></div></form>`)
	return buf.String()
}
