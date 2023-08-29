package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBannerThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M227.38 53.86A4 4 0 0 0 224 52H32a4 4 0 0 0-2.7 7l45 41.29l-45.24 49A4 4 0 0 0 32 156h140l-31.56 66.28a4 4 0 0 0 1.89 5.33a3.92 3.92 0 0 0 1.67.39a4 4 0 0 0 3.61-2.28l80-168a4 4 0 0 0-.23-3.86ZM175.76 148H41.14l41.8-45.29a4 4 0 0 0-.24-5.66L42.28 60h175.39Z"/>`),
		g.Group(children),
	)
}