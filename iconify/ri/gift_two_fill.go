package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.005 13.003v7a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1v-7h16Zm-5.5-11a3.5 3.5 0 0 1 3.163 5.001l3.337-.001a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1v-3a1 1 0 0 1 1-1l3.337.001a3.5 3.5 0 0 1 5.664-3.95a3.48 3.48 0 0 1 2.499-1.051Zm-5 2a1.5 1.5 0 0 0-.145 2.993l.145.007h1.5v-1.5A1.5 1.5 0 0 0 9.649 4.01l-.144-.007Zm5 0l-.145.007a1.5 1.5 0 0 0-1.348 1.348l-.007.145v1.5h1.5l.144-.007a1.5 1.5 0 0 0 0-2.986l-.144-.007Z"/>`),
		g.Group(children),
	)
}