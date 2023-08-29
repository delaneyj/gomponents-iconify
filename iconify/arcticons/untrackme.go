package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Untrackme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.89 15.13l6.59 6.43L25.89 28l-.05-3.66c-5.31.23-6.74 1.22-10.17 4.54c2-10.16 11-10.18 10.15-10Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.07 10.1L26.28 4.92a5.61 5.61 0 0 0-4.26 0L8.89 10.1A2 2 0 0 0 7.62 12v12.47c0 11.16 11.6 19 16.58 19s16.18-7.87 16.18-19V12a2 2 0 0 0-1.31-1.9Z"/>`),
		g.Group(children),
	)
}