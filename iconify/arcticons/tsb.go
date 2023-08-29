package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tsb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="7.688" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.813" cy="24" r="7.688" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.188" cy="24" r="7.688" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.537 20h5.301m-2.65 8v-8m10.32 7.123A2.238 2.238 0 0 0 23.47 28h1.184a1.998 1.998 0 0 0 1.996-2h0a1.998 1.998 0 0 0-1.996-2h-1.308a1.998 1.998 0 0 1-1.996-2h0a1.998 1.998 0 0 1 1.996-2h1.184a2.238 2.238 0 0 1 1.962.876M37.462 24a2 2 0 0 1 0 4h-3.3v-8h3.3a2 2 0 0 1 0 4Zm0 0h-3.299"/>`),
		g.Group(children),
	)
}