package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.144 17.106A9.75 9.75 0 1 0 24 24m2.856 6.894A9.75 9.75 0 1 0 24 24"/><circle cx="14.25" cy="24" r="6.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.75" cy="24" r="6.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.75 26.5l5-5m-5 0l5 5m20.449-4.654l-4.667 4.667l-2.333-2.334"/>`),
		g.Group(children),
	)
}