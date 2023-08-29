package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 20H11v-7h11v7ZM9 20H2v-7h7v7Zm13-9H2V4h20v7Z"/>`),
		g.Group(children),
	)
}