package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#62727a"/><g fill="#fff"><circle cx="32" cy="32" r="4"/><path d="M30 26h4v32h-4z"/><path d="m26.035 35.15l17.102-17.122l2.83 2.827l-17.102 17.121z"/></g><circle cx="32" cy="32" r="3" fill="#62727a"/>`),
		g.Group(children),
	)
}