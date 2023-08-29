package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HundredWords(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.126 23.692V5.196h12.82v16.646m0 4.624v16.338h-12.82V28.546m17.554-7.783V5.196H42.5v13.333m0 4.777v19.498H29.68V25.387M5.5 5.22v19.793m0 4.932v12.858"/>`),
		g.Group(children),
	)
}