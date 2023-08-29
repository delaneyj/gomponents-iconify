package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeniedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m2.5 2.5l10 10m-5 2a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/>`),
		g.Group(children),
	)
}