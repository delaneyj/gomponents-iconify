package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fdroidkdi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.5 9.5h-29a2.08 2.08 0 0 0-2 2v25a2.08 2.08 0 0 0 2 2h37a2.08 2.08 0 0 0 2-2v-17H37a2.81 2.81 0 0 1-2.5-2.5Z"/><circle cx="20" cy="24" r="12" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m44.5 19.5l-10-10"/><circle cx="20" cy="17.74" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20 21.93v9.08"/>`),
		g.Group(children),
	)
}