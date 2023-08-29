package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simplyplural(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.851 42.5h-3.126a6.475 6.475 0 0 1-5.236-2.666l-12.833-17.64c-1.65-2.166-4.55-5.016-4.55-9.24c0-3.987 3.238-7.454 7.743-7.454c4.387 0 7.498 3.467 7.498 7.454c0 4.223-3.421 7.833-9.744 9.032c-7.049 1.338-11.455 5.205-11.455 11.58c0 5.155 3.25 8.934 9.745 8.934c8.54 0 13.565-8.097 20.988-18.048"/>`),
		g.Group(children),
	)
}