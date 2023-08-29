package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboardalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.356 1024h-768q-26 0-45-19t-19-45V128q0-27 19-45.5t45-18.5h128v64h-128v832h768V128h-128V64h128q27 0 45.5 18.5t18.5 45.5v832q0 26-18.5 45t-45.5 19zm-544-576q-13 0-22.5-9.5t-9.5-22.5t9.5-22.5t22.5-9.5h320q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-320zm384 320q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-448q-13 0-22.5-9.5t-9.5-22.5t9.5-22.5t22.5-9.5h448zm-544-160q0-13 9.5-22.5t22.5-9.5h576q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-576q-13 0-22.5-9.5t-9.5-22.5zm448-416h-256q-26 0-45-19t-19-45.5t19-45t45-18.5q0-27 19-45.5t45-18.5h128q27 0 45.5 18.5t18.5 45.5q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19z"/>`),
		g.Group(children),
	)
}