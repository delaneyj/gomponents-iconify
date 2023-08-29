package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryVerticalFullBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M92 12a12 12 0 0 1 12-12h48a12 12 0 0 1 0 24h-48a12 12 0 0 1-12-12Zm112 48v168a28 28 0 0 1-28 28H80a28 28 0 0 1-28-28V60a28 28 0 0 1 28-28h96a28 28 0 0 1 28 28Zm-24 0a4 4 0 0 0-4-4H80a4 4 0 0 0-4 4v168a4 4 0 0 0 4 4h96a4 4 0 0 0 4-4Zm-24 12h-56a12 12 0 0 0 0 24h56a12 12 0 0 0 0-24Zm0 40h-56a12 12 0 0 0 0 24h56a12 12 0 0 0 0-24Zm0 40h-56a12 12 0 0 0 0 24h56a12 12 0 0 0 0-24Zm0 40h-56a12 12 0 0 0 0 24h56a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}