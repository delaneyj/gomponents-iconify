package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ocr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.1 5.5H5.5V18m0 11.9v12.6h11.4m14.1 0h11.5V30m0-11.8V5.5H30.9M9.7 11.2h29.1M9.7 19.5h29.1M9.7 27.8h29.1M9.6 36.2h21.9"/>`),
		g.Group(children),
	)
}