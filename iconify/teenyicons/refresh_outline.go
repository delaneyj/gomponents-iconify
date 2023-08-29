package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 14.5A7 7 0 0 1 3.17 2M7.5.5A7 7 0 0 1 11.83 13m-.33-3v3.5H15M0 1.5h3.5V5"/>`),
		g.Group(children),
	)
}