package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilesBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m220.49 59.51l-40-40A12 12 0 0 0 172 16H92a20 20 0 0 0-20 20v20H56a20 20 0 0 0-20 20v140a20 20 0 0 0 20 20h108a20 20 0 0 0 20-20v-20h20a20 20 0 0 0 20-20V68a12 12 0 0 0-3.51-8.49ZM160 212H60V80h67l33 33Zm40-40h-16v-64a12 12 0 0 0-3.51-8.49l-40-40A12 12 0 0 0 132 56H96V40h71l33 33Zm-56-28a12 12 0 0 1-12 12H88a12 12 0 0 1 0-24h44a12 12 0 0 1 12 12Zm0 40a12 12 0 0 1-12 12H88a12 12 0 0 1 0-24h44a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}