package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopCircleCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M13 8.5a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9ZM6.5 13a6.5 6.5 0 1 1 13 0a6.5 6.5 0 0 1-13 0Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopCircleCircleFilled0)"/></g>`),
		g.Group(children),
	)
}