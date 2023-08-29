package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeSquareF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4zm3.336 13.241H8.34l1.823 1.71a2 2 0 0 0 1.368.54h.205a1.6 1.6 0 0 0 1.6-1.6v-7.8a1.6 1.6 0 0 0-1.6-1.6h-.205a2 2 0 0 0-1.368.541L8.34 6.742H7.336a2 2 0 0 0-2 2v2.5a2 2 0 0 0 2 2zm1.795-4.5l2.205-2.067v6.634L9.13 11.241H7.336v-2.5H9.13z"/>`),
		g.Group(children),
	)
}