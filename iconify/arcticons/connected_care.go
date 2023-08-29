package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectedCare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.497 11.796L24 17.294l-5.497-5.498L24 2.5l5.497 9.296ZM18.196 20.87l-7.51 2.013l-5.306-9.41l10.802-.114l2.014 7.51Zm-2.014 13.771L5.38 34.527l5.305-9.41l7.51 2.013l-2.012 7.511Zm13.315 1.563L24 45.5l-5.497-9.296L24 30.706l5.497 5.498ZM42.62 13.473l-5.305 9.41l-7.51-2.013l2.012-7.511l10.802.114Zm0 21.054l-10.802.114l-2.014-7.51l7.511-2.014l5.305 9.41Z"/>`),
		g.Group(children),
	)
}