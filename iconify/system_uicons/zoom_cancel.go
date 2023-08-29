package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomCancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="translate(3 3)"><circle cx="5.5" cy="5.5" r="5"/><path d="m7.5 7.5l-4-4l4 4zm-4 0l4-4l-4 4zm11 7L9.076 9.076"/></g>`),
		g.Group(children),
	)
}