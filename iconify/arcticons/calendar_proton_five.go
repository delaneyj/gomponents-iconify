package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarProtonFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.885 32.597c1.125.938 2.062 1.313 4.5 1.313h.562a4.964 4.964 0 0 0 4.875-4.875h0a4.964 4.964 0 0 0-4.875-4.875h-5.062v-5.25h9.937"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 25.793v-13.93a4 4 0 0 0-4-4h-31a4 4 0 0 0-4 4v24.275a4 4 0 0 0 4 4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 36.138V11.862a4 4 0 0 0-4-4h-31a4 4 0 0 0-4 4v24.276a4 4 0 0 0 4 4h31a4 4 0 0 0 4-4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.207 40.138V15.819a3.138 3.138 0 0 0-3.138-3.138H4.5"/>`),
		g.Group(children),
	)
}