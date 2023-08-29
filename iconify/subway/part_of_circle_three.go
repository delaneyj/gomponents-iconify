package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartOfCircleThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M221.8 0v256l-181 181c46.4 46.3 110.4 75 181 75c141.4 0 256-114.6 256-256S363.2 0 221.8 0z"/>`),
		g.Group(children),
	)
}