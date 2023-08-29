package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Daedalus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 34.4c.9 0 8.1-4.2 8.1-9.1v-7.7l-8.1-4l-8.1 4v7.7c0 5 7.2 9.1 8.1 9.1Zm0-.2V13.8M16 24h16"/>`),
		g.Group(children),
	)
}