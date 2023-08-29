package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarEighteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.94 37h9m-9-15.525L16.44 19m0 0v18m15.194-9h-2.925a4.513 4.513 0 0 0-4.5 4.5h0a4.513 4.513 0 0 0 4.5 4.5h2.925a4.513 4.513 0 0 0 4.5-4.5h0a4.513 4.513 0 0 0-4.5-4.5Zm0 0a4.513 4.513 0 0 0 4.5-4.5h0a4.513 4.513 0 0 0-4.5-4.5h-2.925a4.513 4.513 0 0 0-4.5 4.5h0a4.513 4.513 0 0 0 4.5 4.5"/><circle cx="32.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}