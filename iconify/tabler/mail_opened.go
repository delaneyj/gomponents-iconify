package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOpened(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 9l9 6l9-6l-9-6l-9 6"/><path d="M21 9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9m0 10l6-6m6 0l6 6"/></g>`),
		g.Group(children),
	)
}