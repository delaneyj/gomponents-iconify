package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryVerticalFullLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M90 8a6 6 0 0 1 6-6h64a6 6 0 0 1 0 12H96a6 6 0 0 1-6-6Zm108 48v168a22 22 0 0 1-22 22H80a22 22 0 0 1-22-22V56a22 22 0 0 1 22-22h96a22 22 0 0 1 22 22Zm-12 0a10 10 0 0 0-10-10H80a10 10 0 0 0-10 10v168a10 10 0 0 0 10 10h96a10 10 0 0 0 10-10Zm-26 18H96a6 6 0 0 0 0 12h64a6 6 0 0 0 0-12Zm0 40H96a6 6 0 0 0 0 12h64a6 6 0 0 0 0-12Zm0 40H96a6 6 0 0 0 0 12h64a6 6 0 0 0 0-12Zm0 40H96a6 6 0 0 0 0 12h64a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}