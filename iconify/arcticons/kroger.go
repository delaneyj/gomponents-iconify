package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kroger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 20.662c0-4.84-3.584-13.245-12.732-13.245s-13.845 7.231-13.845 13.836a11.991 11.991 0 0 1-12.423 12m0-12.453v19.783m13.845 0l-6.324-9.604"/>`),
		g.Group(children),
	)
}