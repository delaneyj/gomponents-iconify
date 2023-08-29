package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crpss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 76.8L435.2 0L256 179.2L76.8 0L0 76.8L179.2 256L0 435.2L76.8 512L256 332.8L435.2 512l76.8-76.8L332.8 256z"/>`),
		g.Group(children),
	)
}