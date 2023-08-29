package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func System(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 19h22V1H1v18Zm4 4h14H5Zm3 0h8v-4H8v4ZM7.757 5.757l2.122 2.122l-2.122-2.122ZM9 10H6h3Zm.879 2.121l-2.122 2.122l2.122-2.122ZM12 13v3v-3Zm2.121-.879l2.122 2.122l-2.122-2.122ZM18 10h-3h3Zm-1.757-4.243l-2.122 2.122l2.122-2.122ZM12 7V4v3Zm0 0a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/>`),
		g.Group(children),
	)
}