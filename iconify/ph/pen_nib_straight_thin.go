package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenNibStraightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M218.75 125.66a1.14 1.14 0 0 0-.1-.19L188 70.94V32a12 12 0 0 0-12-12H80a12 12 0 0 0-12 12v39l-30.66 54.47l-.1.19a12 12 0 0 0 1.15 12.5l86.44 112.28a4 4 0 0 0 6.34 0l86.49-112.35a12 12 0 0 0 1.09-12.43ZM80 28h96a4 4 0 0 1 4 4v36H76V32a4 4 0 0 1 4-4Zm48 120a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm83.27-14.72L132 236.25v-80.59a24 24 0 1 0-8 0v80.58l-79.22-102.9a4 4 0 0 1-.4-4.06L74.35 76h107.31l30 53.28a4 4 0 0 1-.39 4Z"/>`),
		g.Group(children),
	)
}