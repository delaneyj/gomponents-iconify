package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Editor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="31.2" height="39" x="8.4" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.95"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.48 9.37h23.04m-23.04 7.32h23.04M12.48 24h23.04m-23.04 7.31h23.04m-23.04 7.32h23.04"/>`),
		g.Group(children),
	)
}