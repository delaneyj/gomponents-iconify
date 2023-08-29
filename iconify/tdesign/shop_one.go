package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShopOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 7.465V20h4v-7h8v7h4V9.465A3.981 3.981 0 0 1 18 10a3.99 3.99 0 0 1-3-1.354A3.99 3.99 0 0 1 12 10a3.99 3.99 0 0 1-3-1.354A3.99 3.99 0 0 1 6 10a3.982 3.982 0 0 1-2-.535ZM10 6a2 2 0 1 0 4 0V4h-4v2ZM8 4H4v2a2 2 0 1 0 4 0V4Zm8 0v2a2 2 0 1 0 4 0V4h-4Zm-2 16v-5h-4v5h4Z"/>`),
		g.Group(children),
	)
}