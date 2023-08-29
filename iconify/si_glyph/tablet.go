package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.995.043h-8.99A2.005 2.005 0 0 0 2 2.047v11.95C2 15.102 2.899 16 4.005 16h8.99A2.005 2.005 0 0 0 15 13.997V2.047A2.005 2.005 0 0 0 12.995.043zm-2.912 14.999H6.93v-1.114h3.153v1.114zM13 13H4V1.97h9V13z"/>`),
		g.Group(children),
	)
}