package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glaive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M488.7 24.74c-25.6 25.54-51.7 50.93-78 76.26c-57-13.38-129.6 16.5-170.7 43c49.3-8 98.8-16.3 120.4 4.7c-85.4 80.4-173.5 159.8-261.92 239l38.92 44.9c23.9-8.7 56.6-29.2 92-57.6c38-30.6 79.2-70.3 117.4-113.7c67.7-76.8 125.6-166.14 141.9-236.56zM94.96 409.3l-13.61 12.5l19.95 22.6l14-12.9l-20.34-22.2zM68.06 434L18 480.1V494h29.39l40.65-37.4L68.06 434z"/>`),
		g.Group(children),
	)
}