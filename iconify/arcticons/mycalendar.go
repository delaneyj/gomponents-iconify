package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mycalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.5v35a2 2 0 0 0 2 2h2.33v-39H10.4a2 2 0 0 0-2 2Zm4.33-2v39H37.6a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.135 37.557a2.003 2.003 0 0 0 0-.245a2.203 2.203 0 0 0 .5.325a2.325 2.325 0 1 0 2.14-4.127q-.069-.036-.14-.067l-.22-.1l.22-.1a2.74 2.74 0 0 0 .295-.195a2.31 2.31 0 0 0-2.793-3.682l-.002.002a2.003 2.003 0 0 0 0-.245a2.32 2.32 0 0 0-4.635 0a2.003 2.003 0 0 0 0 .245a2.386 2.386 0 0 0-.5-.325a2.326 2.326 0 0 0-2 4.2h0l.22.1l-.22.1a2.678 2.678 0 0 0-.295.19a2.31 2.31 0 1 0 2.795 3.68a2.003 2.003 0 0 0 0 .245a2.32 2.32 0 0 0 4.635 0Zm-2.32-1.905a2.32 2.32 0 1 1 2.32-2.32h0a2.315 2.315 0 0 1-2.31 2.32Z"/>`),
		g.Group(children),
	)
}