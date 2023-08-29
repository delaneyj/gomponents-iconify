package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.488 1024h-768q-26 0-45-19t-19-45V128q0-27 18.5-46t45.5-19h128v128q0 27 19 45.5t45 18.5h384q27 0 45.5-18.5t18.5-45.5V63h128q27 0 45.5 19t18.5 46v832q0 26-18.5 45t-45.5 19zm-256-832h-256q-26 0-45-19t-19-45q0-27 19-45.5t45-18.5q0-27 19-45.5t45-18.5h128q27 0 45.5 18.5t18.5 45.5q27 0 45.5 18.5t18.5 45.5q0 26-18.5 45t-45.5 19z"/>`),
		g.Group(children),
	)
}