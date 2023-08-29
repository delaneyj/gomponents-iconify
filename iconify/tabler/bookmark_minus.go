package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.427 17.256L12 17l-5 3V6a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v9m-1 4h6"/>`),
		g.Group(children),
	)
}