package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ncalc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.232 4.5v14.053m7.027-7.026H26.206M12.838 33.784V17.977a1 1 0 0 1 1.708-.706l16.978 17.053a1 1 0 0 0 1.708-.705V21.895"/><circle cx="12.838" cy="38.642" r="4.858" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}