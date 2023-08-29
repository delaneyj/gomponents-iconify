package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="4" d="M36 27c0 7.217-5.373 17-12 17s-12-9.783-12-17c0-7.216 1.5-11 12-11s12 3.784 12 11Z"/><ellipse stroke="currentColor" stroke-width="4" rx="5" ry="3.5" transform="scale(1 -1) rotate(45 40.625 38.327)"/><ellipse cx="9" cy="17.5" stroke="currentColor" stroke-width="4" rx="5" ry="3.5" transform="rotate(45 9 17.5)"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M12 4c0 6.627 5.373 12 12 12s12-5.373 12-12"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M18 7c0 4.97 2.686 9 6 9s6-4.03 6-9"/><circle cx="20" cy="26" r="2" fill="currentColor"/><circle cx="24" cy="34" r="2" fill="currentColor"/><circle cx="28" cy="26" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}