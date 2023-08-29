package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#0072C6" d="M.5.5h80v200H.5z"/><path fill="#FCD116" d="M75.5.5h225v200h-225z"/><path fill="#009E60" d="M225.5.5h75v200h-75zm-91.666 141.667L150.5 175.5l16.667-33.333l-16.667-33.334zM113 100.5l16.667 33.333l16.667-33.333l-16.667-33.333zm41.667 0l16.667 33.333L188 100.5l-16.667-33.333z"/></g>`),
		g.Group(children),
	)
}