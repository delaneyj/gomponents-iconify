package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmashArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M121.068 15.52v339.31H20.54l100.53 142.516h77.385l-47.49-149.364h50.225V15.52h-80.12zm98.81 0v351.15h-43.362l41.547 130.676h80.36l41.548-130.676h-43.36V15.52h-76.733zm95.42 0v332.462h50.223l-47.487 149.364H391.3L491.823 354.83H391.3V15.52h-76z"/>`),
		g.Group(children),
	)
}