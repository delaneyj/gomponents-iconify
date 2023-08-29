package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.71 16.29l-.15-.12a.76.76 0 0 0-.18-.09L12.2 16a1 1 0 0 0-.91.27a1.15 1.15 0 0 0-.21.33a1 1 0 0 0 1.3 1.31a1.46 1.46 0 0 0 .33-.22a1 1 0 0 0 .21-1.09a1 1 0 0 0-.21-.31ZM16 2H8a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3Zm1 17a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}