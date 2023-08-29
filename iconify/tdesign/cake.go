package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.998 1.998h2.004v2.004H6.998V1.998Zm4 0h2.004v2.004h-2.004V1.998Zm4 0h2.004v2.004h-2.004V1.998ZM9 5v5h2V5h2v5h2V5h2v5h4v10h2v2H1v-2h2V10h4V5h2Zm-4 7v3h1.162L9 15.946l3-1l3 1L17.838 15H19v-3H5Zm14 5h-.838L15 18.054l-3-1l-3 1L5.838 17H5v3h14v-3Z"/>`),
		g.Group(children),
	)
}