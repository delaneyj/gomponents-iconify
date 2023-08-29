package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransgenderOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="256" cy="256" r="128" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="m448 352l-96 96M176 80l-95.98 95.98M464 128V48h-80M48 128V48h80m336 0L346.5 165.5M48 48l117.49 117.49M464 464L346.65 346.37"/>`),
		g.Group(children),
	)
}