package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reallybadchess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.85 18c-.44 5-2.26 6.74-5.39 7.7S11.5 29 11.5 29l-3.08-5.15S15.06 14 22.51 10.37l.43-5l2.62-.91l3.73 4.86h4.47a58.31 58.31 0 0 1 5.82 25.87v6.39a36 36 0 0 1-11.16 1.92a38.11 38.11 0 0 1-11-1.68V33.5l6.76-8.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.58 35.13A36 36 0 0 1 28.42 37a37.91 37.91 0 0 1-11-1.69"/><circle cx="20.97" cy="17.53" r="1.83" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}