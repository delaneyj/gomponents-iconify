package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DuplicateCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilDuplicateCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M18 10.5h-7a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5Zm-6.5 7v-6h6v6h-6Z" clip-rule="evenodd"/><path d="M8.5 14.5h3v1H8a.5.5 0 0 1-.5-.5V8a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 .5.5v3.5h-1v-3h-6v6Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilDuplicateCircleFilled0)"/></g>`),
		g.Group(children),
	)
}