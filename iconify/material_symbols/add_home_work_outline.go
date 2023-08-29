package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddHomeWorkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17v-5h6v5v-5H5v5Zm12-9ZM1 19V9l7-5l7 5v.675q-.575.275-1.075.638t-.925.812V10L8 6.45L3 10v7h2v-5h6v5h.075q.075.525.225 1.025t.375.975H9v-5H7v5H1ZM23 1v10.125q-.425-.45-.925-.813T21 9.675V3h-9v1.4l-2-1.45V1h13Zm-6 6h2V5h-2v2Zm1 14q-2.075 0-3.538-1.463T13 16q0-2.075 1.463-3.538T18 11q2.075 0 3.538 1.463T23 16q0 2.075-1.463 3.538T18 21Zm-.5-2h1v-2.5H21v-1h-2.5V13h-1v2.5H15v1h2.5V19Z"/>`),
		g.Group(children),
	)
}