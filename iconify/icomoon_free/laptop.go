package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 11V3c0-.55-.45-1-1-1H3c-.55 0-1 .45-1 1v8H0v3h16v-3h-2zm-4 2H6v-1h4v1zm3-2H3V3.002L3.002 3h9.996l.002.002V11z"/>`),
		g.Group(children),
	)
}