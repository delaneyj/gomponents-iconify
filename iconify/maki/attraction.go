package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attraction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6 2c-.554 0-.752.505-1 1l-.5 1h-2C1.669 4 1 4.669 1 5.5v5c0 .831.669 1.5 1.5 1.5h10c.831 0 1.5-.669 1.5-1.5v-5c0-.831-.669-1.5-1.5-1.5h-2L10 3c-.25-.5-.446-1-1-1H6zM2.5 5a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zm5 0a3 3 0 1 1 0 6a3 3 0 0 1 0-6zm0 1.5a1.5 1.5 0 0 0 0 3a1.5 1.5 0 0 0 0-3z"/>`),
		g.Group(children),
	)
}