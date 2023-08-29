package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDetails(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1V2.5Zm2 2V19h18V7H10.52l-2-2.5H3Zm3.998 7.498h2.004v2.004H6.998v-2.004Zm4 0h2.004v2.004h-2.004v-2.004Zm4 0h2.004v2.004h-2.004v-2.004Z"/>`),
		g.Group(children),
	)
}