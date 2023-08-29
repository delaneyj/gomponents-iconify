package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fileee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.284 22.437h15.819m-5.995 2.347V43.5l-8.729-18.716M24.658 4.5l-2.55 2.666v15.271M35.716 7.166L33.051 4.5h-6.045m8.711 6.933V9.514"/>`),
		g.Group(children),
	)
}