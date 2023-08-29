package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Digid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.11 35.28v-9.44h2.12A4.13 4.13 0 0 1 16.36 30v1.18a4.13 4.13 0 0 1-4.13 4.13Zm21.53 0v-9.44h2.12A4.13 4.13 0 0 1 37.89 30v1.18a4.13 4.13 0 0 1-4.13 4.13ZM26.26 29v7.08a2.36 2.36 0 0 1-2.36 2.36h0a2.36 2.36 0 0 1-1.67-.69"/><rect width="4.72" height="6.26" x="21.54" y="29.03" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.36" transform="rotate(-180 23.895 32.155)"/><circle cx="18.95" cy="26.13" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.95 29.03v6.25"/><circle cx="29.01" cy="26.13" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.01 29.03v6.25"/>`),
		g.Group(children),
	)
}