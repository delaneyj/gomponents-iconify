package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 14h12v12H18V14Zm0 12h12v12H18V26Zm12-12h12v12H30V14ZM6 26h12v12H6V26Z"/>`),
		g.Group(children),
	)
}