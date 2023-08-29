package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aliexpress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.34 8.27A2.77 2.77 0 0 0 5.5 11.1v28.77a2.8 2.8 0 0 0 2.9 2.63h31.36a2.77 2.77 0 0 0 2.74-3V11a2.8 2.8 0 0 0-3.13-2.68Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.72 8.27h0A2.77 2.77 0 0 0 37 5.5H11a2.77 2.77 0 0 0-2.73 2.77h0"/><circle cx="13.13" cy="22.11" r="1.94" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.86" cy="22.11" r="1.94" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.3 24a10.87 10.87 0 0 0 21.39 0"/>`),
		g.Group(children),
	)
}