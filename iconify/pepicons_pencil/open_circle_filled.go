package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilOpenCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M17.5 17.5v-3.25a.5.5 0 0 1 1 0V18a.5.5 0 0 1-.5.5H8a.5.5 0 0 1-.5-.5V8a.5.5 0 0 1 .5-.5h3.75a.5.5 0 0 1 0 1H8.5v9h9Z"/><path d="M13.354 13.354a.5.5 0 0 1-.708-.708l5-5a.5.5 0 0 1 .708.708l-5 5Z"/><path d="M18.5 11.5a.5.5 0 0 1-1 0V8a.5.5 0 0 1 1 0v3.5Z"/><path d="M14.5 8.5a.5.5 0 0 1 0-1H18a.5.5 0 0 1 0 1h-3.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilOpenCircleFilled0)"/></g>`),
		g.Group(children),
	)
}