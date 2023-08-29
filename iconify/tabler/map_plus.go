package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18.5L9 17l-6 3V7l6-3l6 3l6-3v8.5M9 4v13m6-10v8m1 4h6m-3-3v6"/>`),
		g.Group(children),
	)
}