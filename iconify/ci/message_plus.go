package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessagePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H9c-.433 0-.854.14-1.2.4L3 21ZM5 5v12l2.134-1.6a1.99 1.99 0 0 1 1.2-.4H19V5H5Zm8 9h-2v-3H8V9h3V6h2v3h3v2h-3v3Z"/>`),
		g.Group(children),
	)
}