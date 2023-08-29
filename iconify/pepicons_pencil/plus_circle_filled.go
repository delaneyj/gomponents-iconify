package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPlusCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M8 13.5a.5.5 0 0 1 0-1h10a.5.5 0 0 1 0 1H8Z"/><path d="M12.5 8a.5.5 0 0 1 1 0v10a.5.5 0 0 1-1 0V8Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPlusCircleFilled0)"/></g>`),
		g.Group(children),
	)
}