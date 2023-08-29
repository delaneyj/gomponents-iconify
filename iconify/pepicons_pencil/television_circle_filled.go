package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTelevisionCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M6.25 12v5a3 3 0 0 0 3 3h7.5a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3h-7.5a3 3 0 0 0-3 3Zm3 7a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h7.5a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-7.5Z" clip-rule="evenodd"/><path d="m12.86 8.688l2-2.5c.416-.52 1.197.104.78.624l-2 2.5c-.416.52-1.197-.104-.78-.624Z"/><path d="m11.86 9.312l-2-2.5c-.417-.52.364-1.145.78-.624l2 2.5c.417.52-.364 1.145-.78.624Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTelevisionCircleFilled0)"/></g>`),
		g.Group(children),
	)
}