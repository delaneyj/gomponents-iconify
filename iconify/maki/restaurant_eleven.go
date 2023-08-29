package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestaurantEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M2.25 0l-.5 3.5c-.07.495 1.24 1.4 1.25 2V11h1V5.5c0-.6 1.32-1.505 1.25-2L4.75 0h-.5l.25 2.75l-.75.5V0h-.5v3.25l-.75-.5L2.75 0h-.5zM9 0C7.5 0 7.006 1.724 7 3v3h1v5h1V0z" fill="currentColor"/>`),
		g.Group(children),
	)
}