package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceAwayTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 11.31V6.5a1.5 1.5 0 0 0-3 0V12c0 .438.191.854.524 1.139l3.5 3a1.5 1.5 0 0 0 1.952-2.278L13 11.311ZM0 12C0 5.373 5.373 0 12 0s12 5.373 12 12s-5.373 12-12 12S0 18.627 0 12Zm12-9a9 9 0 1 0 0 18a9 9 0 0 0 0-18Z"/>`),
		g.Group(children),
	)
}