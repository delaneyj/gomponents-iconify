package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelationalGraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRelationalGraph0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 35h40"/><path d="M24 32c0-4.418-3.806-8-8.5-8S7 27.582 7 32"/><path stroke-linecap="round" stroke-linejoin="round" d="M41 32c0-11.046-7.611-20-17-20S7 20.954 7 32"/><circle cx="41" cy="35" r="3" fill="#555" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="35" r="3" fill="#555" stroke-linecap="round" stroke-linejoin="round"/><circle cx="7" cy="35" r="3" fill="#555" stroke-linecap="round" stroke-linejoin="round"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRelationalGraph0)"/>`),
		g.Group(children),
	)
}