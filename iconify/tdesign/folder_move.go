package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderMove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1V2.5Zm2 2V19h18V7H10.52l-2-2.5H3Z"/><path fill="currentColor" d="M8 11.999h4.657l-1.466-1.53l1.445-1.384l3.749 3.914l-3.75 3.914l-1.444-1.384L12.657 14H8v-2Z"/>`),
		g.Group(children),
	)
}