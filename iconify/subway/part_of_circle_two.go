package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartOfCircleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M119.5 0v299.9L331.6 512c54.3-54.3 87.8-129.2 87.8-212.1C419.5 134.3 285.2 0 119.5 0z"/>`),
		g.Group(children),
	)
}