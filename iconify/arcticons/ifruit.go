package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ifruit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 24.082a18.5 18.5 0 0 0 37 0Zm20.063 0a8.381 8.381 0 0 0-16.762 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.2 24.082a8.382 8.382 0 0 0-15.207-4.866m-7.908-3.442c.63-4.775 12.333-13.82 21.02-8.966v2.045s-9.547.815-14.418 8.91"/>`),
		g.Group(children),
	)
}