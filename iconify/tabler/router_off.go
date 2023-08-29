package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RouterOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 13h2a2 2 0 0 1 2 2v2m-.588 3.417c-.362.36-.861.583-1.412.583H5a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h8m4 4v.01M13 17v.01m-.774-8.81a4 4 0 0 1 6.024.55M9.445 5.407A8 8 0 0 1 21.5 6.5M3 3l18 18"/>`),
		g.Group(children),
	)
}