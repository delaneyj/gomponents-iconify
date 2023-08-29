package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CountdownCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCountdownCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M13 6a.5.5 0 0 1 .5-.5A7.5 7.5 0 1 1 6 13a.5.5 0 0 1 1 0a6.5 6.5 0 1 0 6.5-6.5a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path d="M11 6.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm-2.5 1.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0ZM7 10.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/><path fill-rule="evenodd" d="M8.947 14.224a.5.5 0 0 0-.223-.671l-2-1a.5.5 0 0 0-.448.894l2 1a.5.5 0 0 0 .671-.223Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.854 12.646a.5.5 0 0 1 0 .708l-1.5 1.5a.5.5 0 1 1-.708-.707l1.5-1.5a.5.5 0 0 1 .708 0ZM13.5 9.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-1 0v-3a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 13a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1 0-1h3a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCountdownCircleFilled0)"/></g>`),
		g.Group(children),
	)
}