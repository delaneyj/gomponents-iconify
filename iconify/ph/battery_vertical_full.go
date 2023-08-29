package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryVerticalFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 8a8 8 0 0 1 8-8h64a8 8 0 0 1 0 16H96a8 8 0 0 1-8-8Zm112 48v168a24 24 0 0 1-24 24H80a24 24 0 0 1-24-24V56a24 24 0 0 1 24-24h96a24 24 0 0 1 24 24Zm-16 0a8 8 0 0 0-8-8H80a8 8 0 0 0-8 8v168a8 8 0 0 0 8 8h96a8 8 0 0 0 8-8Zm-24 16H96a8 8 0 0 0 0 16h64a8 8 0 0 0 0-16Zm0 40H96a8 8 0 0 0 0 16h64a8 8 0 0 0 0-16Zm0 40H96a8 8 0 0 0 0 16h64a8 8 0 0 0 0-16Zm0 40H96a8 8 0 0 0 0 16h64a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}