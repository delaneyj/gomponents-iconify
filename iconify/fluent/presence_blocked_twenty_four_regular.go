package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceBlockedTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 12c0 6.627-5.373 12-12 12S0 18.627 0 12S5.373 0 12 0s12 5.373 12 12Zm-3 0a8.959 8.959 0 0 0-1.664-5.215L6.786 19.336A9 9 0 0 0 21 12Zm-3.785-7.336a9 9 0 0 0-12.55 12.55l12.55-12.55Z"/>`),
		g.Group(children),
	)
}