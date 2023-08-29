package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equals(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.585 768h-896q-27 0-45.5-19t-18.5-45V576q0-27 18.5-45.5t45.5-18.5h896q27 0 45.5 18.5t18.5 45.5v128q0 27-18.5 45.5t-45.5 18.5zm0-512h-896q-27 0-45.5-19t-18.5-45V64q0-27 18.5-45.5T64.585 0h896q27 0 45.5 18.5t18.5 45.5v128q0 26-18.5 45t-45.5 19z"/>`),
		g.Group(children),
	)
}