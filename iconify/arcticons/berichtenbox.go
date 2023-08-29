package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Berichtenbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.47 10.71a2 2 0 0 0-2 2h0v22.61a2 2 0 0 0 2 2h35.06a2 2 0 0 0 2-2h0V12.68a2 2 0 0 0-2-2H6.47Zm33.21 3.82L24 26.07L8.32 14.53m17.96-3.85h-4.56v8.554h4.56Zm-9.803 15.98l-8.609 7.643m32.264 0l-8.61-7.642"/>`),
		g.Group(children),
	)
}