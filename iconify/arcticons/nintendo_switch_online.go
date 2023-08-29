package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NintendoSwitchOnline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="13.5" cy="16.2" r="3.4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.5" cy="26.2" r="3.4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.5 5.5h-7.3c-4.8 0-8.7 3.9-8.7 8.6V34c0 4.7 3.9 8.6 8.7 8.6h7.3V5.5Zm12.3 0h-7.3v37.1h7.3c4.8 0 8.7-3.9 8.7-8.6V14.1c0-4.7-3.9-8.6-8.7-8.6Z"/>`),
		g.Group(children),
	)
}