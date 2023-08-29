package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HideMyApplist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 43.5h27a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2h-27a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2Zm17-39v39m-7-19.5h7"/><circle cx="16" cy="24" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.5 14h7"/><circle cx="16" cy="14" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.5 34h7"/><circle cx="16" cy="34" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}