package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Easel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M243.5 22.5v52.8h25V22.5h-25zm-14.6 67.1v19.5h54.2V89.6h-54.2zm-57.7 34.1v202.2h169.4V123.7H171.2zm-21.6 219.6v21h212.8v-21H149.6zm21 35.3L144.2 485l24.3 6l27.8-112.4h-25.7zm145.1 0L343.5 491l24.3-6l-26.4-106.4h-25.7zm-67 2.5v70h16.7v-70h-16.7z"/>`),
		g.Group(children),
	)
}