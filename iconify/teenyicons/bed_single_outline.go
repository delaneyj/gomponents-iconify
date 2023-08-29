package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedSingleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M14.5 4v11M.5 0v15M0 10.5h15m-15-3h15m-13-2h4"/>`),
		g.Group(children),
	)
}