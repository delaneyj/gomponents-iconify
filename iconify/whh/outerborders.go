package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outerborders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M0 960V0h960v960H0zm896-448h-64v-64h64V64H512v64h-64V64H64v384h64v64H64v384h384v-64h64v64h384V512zm-256 0h-64v-64h64v64zm128 0h-64v-64h64v64zM448 704h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64v-64zM320 448h64v64h-64v-64zm-128 0h64v64h-64v-64z"/>`),
		g.Group(children),
	)
}