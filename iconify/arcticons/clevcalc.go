package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clevcalc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5h33c1.1 0 2 .9 2 2v33c0 1.1-.9 2-2 2h-33c-1.1 0-2-.9-2-2v-33c0-1.1.9-2 2-2Zm21 10.5H37M12.1 29.1l7.6 7.5m-.1-7.5L12 36.7m3.9-25v8.5M11.6 16h8.5m9.4 15H36m-6.5 3.6H36"/><circle cx="32.7" cy="32.8" r="6.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}