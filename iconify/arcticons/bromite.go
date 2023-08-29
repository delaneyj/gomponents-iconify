package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bromite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.17 12.88a1.41 1.41 0 0 0 1.77-.35l1.18-1.85a1.41 1.41 0 0 0-.33-1.77l-9.7-6.24a1.41 1.41 0 0 0-1.77.33l-1.18 1.86a1.43 1.43 0 0 0 .33 1.77l.85.6l-1.69 2.69a17.73 17.73 0 0 0-4.81-.67a18.13 18.13 0 1 0 12.83 5.31l1.51-2.29Zm-.84 14.49a14.57 14.57 0 0 1-.6 4.15a12.21 12.21 0 0 1-5.22 1.17a12.92 12.92 0 0 1-8.26-3.21a21.89 21.89 0 0 0-3.91-2.69a21.91 21.91 0 0 0-10.42-2.87a14.41 14.41 0 0 1 28.41 3.45Z"/>`),
		g.Group(children),
	)
}