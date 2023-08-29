package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChipExtraction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-1.875 0-3.513-.713t-2.85-1.924q-1.212-1.213-1.924-2.85T3 12q0-1.875.713-3.513t1.924-2.85q1.213-1.212 2.85-1.924T12 3v2Q9.075 5 7.037 7.038T5 12q0 2.925 2.038 4.963T12 19v2Zm4-4l-1.4-1.425L17.175 13H9v-2h8.175L14.6 8.4L16 7l5 5l-5 5Z"/>`),
		g.Group(children),
	)
}