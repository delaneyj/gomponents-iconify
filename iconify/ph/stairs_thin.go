package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StairsThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 28H56a12 12 0 0 0-12 12v176a12 12 0 0 0 12 12h144a12 12 0 0 0 12-12V40a12 12 0 0 0-12-12Zm-48 112h52v32h-96v-32Zm4-8v-32h48v32ZM56 36h144a4 4 0 0 1 4 4v52h-52a4 4 0 0 0-4 4v36h-44a4 4 0 0 0-4 4v36H52V40a4 4 0 0 1 4-4Zm144 184H56a4 4 0 0 1-4-4v-36h152v36a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}