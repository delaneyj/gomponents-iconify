package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LenovoLife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 12.988c-2.92-1.751-7.074-3.032-12.974-3.032c-18.657 0-27.521 11.338-25.82 24.274c4.097 3.247 7.498 3.814 13.554 3.814c6.709 0 24.558-3.86 25.24-25.056Z"/>`),
		g.Group(children),
	)
}