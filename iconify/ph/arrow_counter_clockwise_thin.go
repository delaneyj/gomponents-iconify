package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCounterClockwiseThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 128a92 92 0 0 1-90.77 92H128a91.47 91.47 0 0 1-63.13-25.1a4 4 0 1 1 5.5-5.82A84 84 0 1 0 68.6 68.57l-.13.12L34.3 100H72a4 4 0 0 1 0 8H24a4 4 0 0 1-4-4V56a4 4 0 0 1 8 0v38.89l35-32A92 92 0 0 1 220 128Z"/>`),
		g.Group(children),
	)
}