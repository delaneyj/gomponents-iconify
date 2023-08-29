package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 44H32a12 12 0 0 0-12 12v144a12 12 0 0 0 12 12h192a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12ZM92 140v-24h72v24Zm-48 0v-24h24v24Zm144-24h24v24h-24Zm24-24h-72V68h72Zm-96-24v24H44V68Zm-72 96h72v24H44Zm96 24v-24h72v24Z"/>`),
		g.Group(children),
	)
}