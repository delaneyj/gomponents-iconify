package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tubemate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.72 4.5c-1.35 5.23 3.21 15.32 9.41 15.91A4.12 4.12 0 0 0 21.86 17s10.08 3 9.94 16a15.12 15.12 0 0 0 7.2-3.25c.3.15-2.17 12-9.19 13.75c0 0-12.18-1.94-14.72-8c0 0 5 1.34 7.92 0c-.26-2.58-5.14-5.2-10-9a9.7 9.7 0 0 1-4-7.25l3.17 2.06S8 10.85 13.72 4.5Z"/>`),
		g.Group(children),
	)
}