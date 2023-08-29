package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunHorizonBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 148h-36.11c.07-1.33.11-2.66.11-4a76 76 0 0 0-152 0c0 1.34 0 2.67.11 4H16a12 12 0 0 0 0 24h224a12 12 0 0 0 0-24Zm-164-4a52 52 0 0 1 104 0c0 1.34-.07 2.67-.17 4H76.17c-.1-1.33-.17-2.66-.17-4Zm144 56a12 12 0 0 1-12 12H48a12 12 0 0 1 0-24h160a12 12 0 0 1 12 12ZM12.62 92.21a12 12 0 0 1 15.17-7.59l12 4a12 12 0 1 1-7.58 22.77l-12-4a12 12 0 0 1-7.59-15.18Zm56-48.41a12 12 0 1 1 22.76-7.59l4 12a12 12 0 1 1-22.76 7.59Zm140 60a12 12 0 0 1 7.59-15.18l12-4a12 12 0 0 1 7.58 22.77l-12 4a12 12 0 0 1-15.17-7.59Zm-48-55.59l4-12a12 12 0 1 1 22.76 7.59l-4 12a12 12 0 1 1-22.76-7.59Z"/>`),
		g.Group(children),
	)
}