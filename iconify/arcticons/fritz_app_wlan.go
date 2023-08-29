package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FritzAppWlan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.582 13.344l.009-8.232M8.25 13.857V4.5h4.678M8.25 9.178h3.041m3.374 5.849V7.109h2.574a2.672 2.672 0 1 1 0 5.345h-2.574m13.613 2.473V6.83m-2.384-.09h4.768m1.305-1.8h4.769l-4.769 8.098h4.769M19.91 15.027l-2.483-2.582M39.615 5.39l-.45 4.499"/><circle cx="39.03" cy="11.688" r=".72" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.731 34.158a7.712 7.712 0 0 0-9.854-.027v.027m4.908 3.157a3.092 3.092 0 1 0 3.092 3.093v0h0a3.092 3.092 0 0 0-3.092-3.093Zm9.649-8.504a15.164 15.164 0 0 0-19.3 0M39.071 23.2a22.419 22.419 0 0 0-28.603 0"/>`),
		g.Group(children),
	)
}