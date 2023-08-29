package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperFrog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M248.2 77.85L243.1 99l30.9-3.3l-25.8-17.85zm38 32.55l-35 1.5l40 22.6l-5-24.1zm-49.9 16.8l-17.5 15.1l165.8 98.4l-15.1 51.5L421 234L236.3 127.2zM179.5 148L73.1 230l6.99 152L260 355.4l94.8-109.3L179.5 148zm206.3 41.2l39.8 25.8l-2.1-24.1l-37.7-1.7zm56.6 7l3.8 36.6l-12 13.5l65.5-17.6l-57.3-32.5zM58.83 222.7l-37.66 68.5l32.87-9.4l4.79-59.1zM314.7 324.6l-31.1 32.1l37-7.4l-5.9-24.7zm-1.4 46.1L245 388l2.3 44.5l66-61.8z"/>`),
		g.Group(children),
	)
}