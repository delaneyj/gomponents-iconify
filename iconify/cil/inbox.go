package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 16v480h448V16Zm416 288H342.111l-32 64H201.889l-32-64H64V144h384Zm0-256v64H64V48ZM64 464V336h86.111l32 64h147.778l32-64H448v128Z"/>`),
		g.Group(children),
	)
}