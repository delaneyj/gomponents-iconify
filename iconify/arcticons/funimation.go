package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Funimation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.238 31.34c-.615 2.436-1.974 4.155-4.228 5.044c-2.623 1.036-5.305 1.025-7.931-.002c-1.95-.762-3.26-2.204-3.964-4.181c-.09-.253-.136-.521-.222-.86h16.345Z"/>`),
		g.Group(children),
	)
}