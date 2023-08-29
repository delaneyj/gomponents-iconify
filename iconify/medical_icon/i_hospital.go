package medical_icon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IHospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M20.13 34.78v27.8H7.18V.592h12.95v23.483h24.173V.592h12.95V62.58h-12.95v-27.8H20.13z"/>`),
		g.Group(children),
	)
}