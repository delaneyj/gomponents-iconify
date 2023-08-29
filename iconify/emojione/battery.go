package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Battery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="#3e4347"><path d="M42 7c0 1.1-.9 2-2 2H24c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h16c1.1 0 2 .9 2 2v3"/><path d="M48 6H16c-2.2 0-4 1.8-4 4v48c0 2.2 1.8 4 4 4h32c2.2 0 4-1.8 4-4V10c0-2.2-1.8-4-4-4m0 47c0 .5-.5 1-1 1H17c-.5 0-1-.5-1-1V15c0-.5.5-1 1-1h30c.5 0 1 .5 1 1v38"/></g><path fill="#a8d600" d="M46 24c0 1.1-.9 2-2 2H20c-1.1 0-2-.9-2-2v-4c0-1.1.9-2 2-2h24c1.1 0 2 .9 2 2v4m0 12c0 1.1-.9 2-2 2H20c-1.1 0-2-.9-2-2v-4c0-1.1.9-2 2-2h24c1.1 0 2 .9 2 2v4m0 12c0 1.1-.9 2-2 2H20c-1.1 0-2-.9-2-2v-4c0-1.1.9-2 2-2h24c1.1 0 2 .9 2 2v4"/>`),
		g.Group(children),
	)
}