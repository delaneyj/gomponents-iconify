package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picardbarcodescanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.291 41.805A21.497 21.497 0 1 1 42.98 33.87"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.711 23.01a7.329 7.329 0 0 1 14.59.127"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.642 23.124a3.42 3.42 0 0 1 6.715 0"/><rect width="30.291" height="22.474" x="12.557" y="23.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.954"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.182 25.73v17.59m1.661-17.59v17.59m4.887-17.59v17.59m5.178-17.59v17.59m1.954-17.59v17.59m3.908-17.59v17.59m1.955-17.59v17.59m3.909-17.59v17.59M26.345 25.73v17.59"/>`),
		g.Group(children),
	)
}