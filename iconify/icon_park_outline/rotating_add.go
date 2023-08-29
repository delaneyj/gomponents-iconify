package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotatingAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6.676 14C10.134 8.022 16.597 4 24 4M6.676 14H14m-7.324 0V6.676M14 41.324C8.022 37.866 4 31.403 4 24m10 17.324V34m0 7.324H6.676M41.324 34C37.866 39.978 31.403 44 24 44m17.324-10H34m7.324 0v7.324M34 6.676C39.978 10.134 44 16.597 44 24M34 6.676V14m0-7.324h7.324M18 24h12m-6 6V18"/>`),
		g.Group(children),
	)
}