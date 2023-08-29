package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Milkshake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 10a5 5 0 0 0-10 0m-1 1a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1zm1 2l1.81 7.243a1 1 0 0 0 .97.757h4.44a1 1 0 0 0 .97-.757L17 13m-5-8V3"/>`),
		g.Group(children),
	)
}