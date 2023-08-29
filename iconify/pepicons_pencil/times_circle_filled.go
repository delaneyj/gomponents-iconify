package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimesCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTimesCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M9.854 16.854a.5.5 0 0 1-.708-.708l7-7a.5.5 0 0 1 .708.708l-7 7Z"/><path d="M9.146 9.854a.5.5 0 1 1 .708-.708l7 7a.5.5 0 0 1-.708.708l-7-7Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTimesCircleFilled0)"/></g>`),
		g.Group(children),
	)
}