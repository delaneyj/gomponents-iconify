package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7h12a1 1 0 0 1 1 1v12M7 10v6a1 1 0 0 0 1 1h6M7 4v3m10 10h3"/>`),
		g.Group(children),
	)
}