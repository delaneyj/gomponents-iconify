package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudLightningLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M156 18a74.21 74.21 0 0 0-69.89 49.3a6.18 6.18 0 0 0-1.83-.62A50.83 50.83 0 0 0 76 66a50 50 0 0 0 0 100h41.4l-18.55 30.91A6 6 0 0 0 104 206h29.4l-18.55 30.91a6 6 0 0 0 10.3 6.18l24-40A6 6 0 0 0 144 194h-29.4l16.8-28H156a74 74 0 0 0 0-148Zm0 136H76a38 38 0 1 1 6.31-75.48a6.82 6.82 0 0 0 .79.08a72.86 72.86 0 0 0-1.1 9.05a6 6 0 0 0 12 .7A62.06 62.06 0 1 1 156 154Z"/>`),
		g.Group(children),
	)
}