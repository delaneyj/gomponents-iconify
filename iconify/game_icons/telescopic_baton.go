package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelescopicBaton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M327.672 195.564L438.26 84.986a25.535 25.535 0 1 0-11.828-11.922l-110.683 110.63zM151.888 336.86l23.75 23.75L65.66 470.581a16.794 16.794 0 0 1-23.75-23.751zm149.808-143.522l16.33 16.33l-134.79 134.79l-16.33-16.33z"/>`),
		g.Group(children),
	)
}