package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HulfOfCircleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M159 0v512c141.4 0 256-114.6 256-256S300.4 0 159 0z"/>`),
		g.Group(children),
	)
}