package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleBigCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopCircleBigCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M13 5.5a7.5 7.5 0 1 0 0 15a7.5 7.5 0 0 0 0-15ZM3.5 13a9.5 9.5 0 1 1 19 0a9.5 9.5 0 0 1-19 0Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopCircleBigCircleFilled0)"/></g>`),
		g.Group(children),
	)
}