package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tennis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 4C23.8991 10.6682 22.2619 15.6696 19.0884 19.0044C15.9148 22.3391 10.8853 24.0071 4 24.0083"/><path stroke-linecap="round" d="M43.9682 25.0052C37.4557 24.5585 32.4795 25.9505 29.0395 29.1812C25.5994 32.4119 23.9206 37.3515 24.0029 43.9999"/></g>`),
		g.Group(children),
	)
}