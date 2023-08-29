package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarProtonTwenty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.007 28.91a4.98 4.98 0 0 0 5 5a4.825 4.825 0 0 0 4.815-5v-5a4.943 4.943 0 0 0-4.815-5a5.1 5.1 0 0 0-5 5Zm-12.122-5a4.943 4.943 0 0 1 4.815-5a4.976 4.976 0 0 1 3.518 8.518c-2.037 1.667-8.333 6.482-8.333 6.482H18.7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 25.793v-13.93a4 4 0 0 0-4-4h-31a4 4 0 0 0-4 4v24.275a4 4 0 0 0 4 4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 36.138V11.862a4 4 0 0 0-4-4h-31a4 4 0 0 0-4 4v24.276a4 4 0 0 0 4 4h31a4 4 0 0 0 4-4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.207 40.138V15.819a3.138 3.138 0 0 0-3.138-3.138H4.5"/>`),
		g.Group(children),
	)
}