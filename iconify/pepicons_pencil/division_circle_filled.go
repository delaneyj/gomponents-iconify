package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DivisionCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilDivisionCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M8 13.5a.5.5 0 0 1 0-1h10a.5.5 0 0 1 0 1H8Z"/><path fill-rule="evenodd" d="M11.5 9a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Zm2 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-2 8a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Zm2 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilDivisionCircleFilled0)"/></g>`),
		g.Group(children),
	)
}