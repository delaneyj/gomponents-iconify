package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 7h-7L10 .5L7 7H0l5.46 5.47l-1.64 7l6.18-3.7l6.18 3.73l-1.63-7z"/>`),
		g.Group(children),
	)
}