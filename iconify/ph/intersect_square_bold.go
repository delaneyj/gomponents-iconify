package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntersectSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 84h-44V40a12 12 0 0 0-12-12H40a12 12 0 0 0-12 12v120a12 12 0 0 0 12 12h44v44a12 12 0 0 0 12 12h120a12 12 0 0 0 12-12V96a12 12 0 0 0-12-12ZM52 148V52h96v32H96a12 12 0 0 0-12 12v52Zm56-23l23 23h-23Zm40 6l-23-23h23Zm56 73h-96v-32h52a12 12 0 0 0 12-12v-52h32Z"/>`),
		g.Group(children),
	)
}