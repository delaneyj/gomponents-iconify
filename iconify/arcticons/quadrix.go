package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quadrix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.39 14.39h19.22v19.22H14.39z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.74 44.722c-3.532-.893-7.552-2.256-10.07-2.256l-24.12-.006a5 5 0 0 1-4.998-5V10.535a5 5 0 0 1 5-5h26.55a5.362 5.362 0 0 1 5.375 5.374l.025 24.292c0 2.412 1.317 6.206 2.238 9.522Z"/>`),
		g.Group(children),
	)
}