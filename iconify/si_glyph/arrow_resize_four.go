package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowResizeFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.12 16.003L16 14.122l-3.038-3.057L15.018 9H9v6.047l2.087-2.097l3.033 3.053zM7 .969L4.913 3.065L1.88.013L0 1.894L3.038 4.95L1.042 6.917H7V.969z"/>`),
		g.Group(children),
	)
}