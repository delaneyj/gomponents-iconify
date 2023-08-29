package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Satellite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 7L9 3L5 7l4 4m8 0l4 4l-4 4l-4-4"/><path d="m8 12l4 4l6-6l-4-4Zm8-4l3-3M9 21a6 6 0 0 0-6-6"/></g>`),
		g.Group(children),
	)
}