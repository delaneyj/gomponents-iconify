package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUpOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11v8a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h3a3.987 3.987 0 0 0 2.828-1.172M11 7V6a2 2 0 1 1 4 0v5h3a2 2 0 0 1 2 2c-.222 1.112-.39 1.947-.5 2.503m-.758 3.244C18.35 19.57 17.698 20.059 17 20h-7a3 3 0 0 1-3-3M3 3l18 18"/>`),
		g.Group(children),
	)
}