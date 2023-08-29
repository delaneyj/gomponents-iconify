package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPersonCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g stroke="#000" stroke-linecap="round" transform="translate(3 3)"><circle cx="9.5" cy="5.5" r="3"/><path d="M15 16.5v-2c0-3.098-2.495-6-5.5-6c-3.006 0-5.5 2.902-5.5 6v2"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPersonCircleFilled0)"/></g>`),
		g.Group(children),
	)
}