package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarProtonEighteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.179 20.972l3.75-2.062m0 0v15m9.849-7.5h-2.437a3.761 3.761 0 0 0-3.75 3.75h0a3.761 3.761 0 0 0 3.75 3.75h2.437a3.761 3.761 0 0 0 3.75-3.75h0a3.761 3.761 0 0 0-3.75-3.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.778 26.41a3.761 3.761 0 0 0 3.75-3.75h0a3.761 3.761 0 0 0-3.75-3.75h-2.437a3.761 3.761 0 0 0-3.75 3.75h0a3.761 3.761 0 0 0 3.75 3.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 25.793v-13.93a4 4 0 0 0-4-4h-31a4 4 0 0 0-4 4v24.275a4 4 0 0 0 4 4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 36.138V11.862a4 4 0 0 0-4-4h-31a4 4 0 0 0-4 4v24.276a4 4 0 0 0 4 4h31a4 4 0 0 0 4-4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.207 40.138V15.819a3.138 3.138 0 0 0-3.138-3.138H4.5"/>`),
		g.Group(children),
	)
}