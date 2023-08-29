package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NinetyNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.52" cy="21.62" r="4.76" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.28 29.38a4.35 4.35 0 0 0 3.88 1.76h.36a4.75 4.75 0 0 0 4.76-4.76v-4.76"/><circle cx="30.48" cy="21.62" r="4.76" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.25 29.38a4.35 4.35 0 0 0 3.88 1.76h.35a4.74 4.74 0 0 0 4.76-4.76v-4.76"/>`),
		g.Group(children),
	)
}