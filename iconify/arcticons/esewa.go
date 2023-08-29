package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Esewa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.69 27.75a7.91 7.91 0 0 1-5.1 6.12h0a7.94 7.94 0 0 1-10.16-4.74l-1.76-4.84a7.94 7.94 0 0 1 4.74-10.16h0a7.94 7.94 0 0 1 10.16 4.74l.88 2.42l-14.9 5.42M45.5 24h-8"/>`),
		g.Group(children),
	)
}