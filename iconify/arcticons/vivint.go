package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vivint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.55 8.83l-12 10.92a3.24 3.24 0 0 0-1.06 2.4v16.22a1.08 1.08 0 0 0 1.08 1.08H31a1.08 1.08 0 0 0 1.08-1.08V22.15a3.24 3.24 0 0 0-1.08-2.4L19 8.83a1.09 1.09 0 0 0-1.45 0Z"/><circle cx="39.18" cy="35.12" r="4.32" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}