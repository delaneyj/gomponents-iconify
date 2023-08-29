package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoHands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m37 19.976l-.983 3.957l-8.706 6.307l.05 11.718L34 42l-.373-8.485c6.922-4.33 10.383-7.5 10.383-9.515V6.06M11 20.051l1.033 3.95l8.368 6.415l.359 11.712L14 42l.203-8.091C7.409 30 4.013 27.025 4.013 24.98V6.03"/>`),
		g.Group(children),
	)
}