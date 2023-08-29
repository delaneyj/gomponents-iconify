package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v7.5h-2V7H10.52l-2-2.5H3V19h8.5v2H1V2.5Zm16.25 12a2.75 2.75 0 0 1 1.946 4.693l-.008.008A2.75 2.75 0 1 1 17.25 14.5Zm3.992 5.325a4.75 4.75 0 1 0-1.414 1.415l1.67 1.674l1.416-1.412l-1.672-1.677Z"/>`),
		g.Group(children),
	)
}