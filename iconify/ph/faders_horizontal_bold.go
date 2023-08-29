package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FadersHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M28 80a12 12 0 0 1 12-12h28a12 12 0 0 1 0 24H40a12 12 0 0 1-12-12Zm188 84h-32v-12a12 12 0 0 0-24 0v48a12 12 0 0 0 24 0v-12h32a12 12 0 0 0 0-24Zm-84 0H40a12 12 0 0 0 0 24h92a12 12 0 0 0 0-24Zm-24-48a12 12 0 0 0 12-12V92h96a12 12 0 0 0 0-24h-96V56a12 12 0 0 0-24 0v48a12 12 0 0 0 12 12Z"/>`),
		g.Group(children),
	)
}