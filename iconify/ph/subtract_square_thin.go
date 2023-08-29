package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 92h-52V40a4 4 0 0 0-4-4H40a4 4 0 0 0-4 4v120a4 4 0 0 0 4 4h52v52a4 4 0 0 0 4 4h120a4 4 0 0 0 4-4V96a4 4 0 0 0-4-4Zm-54.34 120l-48-48h44.68l48 48Zm2.34-53.66v-44.68l48 48v44.68ZM212 100v50.34l-48-48V100ZM44 44h112v112H44Zm56 120h2.34l48 48H100Z"/>`),
		g.Group(children),
	)
}