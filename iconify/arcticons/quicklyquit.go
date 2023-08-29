package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quicklyquit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.63 4.5A3.66 3.66 0 1 1 25 8.16a3.66 3.66 0 0 1 3.63-3.66Zm10 20.87a13.29 13.29 0 0 1-10-4.54l-1.1 5.44l3.82 3.65V43.5h-3.66V32.62L23.89 29l-1.82 8l-12.69-2.57l.71-3.65L19 32.59l2.93-14.68l-3.27 1.21v6.17H15v-8.53l9.44-4c.55 0 .91-.19 1.45-.19A3.73 3.73 0 0 1 29 14.41l1.82 2.9a8.91 8.91 0 0 0 7.84 4.42Z"/>`),
		g.Group(children),
	)
}