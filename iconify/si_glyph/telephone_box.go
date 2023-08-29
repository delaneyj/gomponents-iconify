package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelephoneBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 15.021V4h-1V2.917h.852C12.852 1.306 11.021.083 8 .083c-3.021 0-5 1.223-5 2.834h1.023V4H3v11.021H2V16h12v-.979h-1zM5 3h6v1H5V3zm1 12.021H5v-2.04h1v2.04zM6 12H5v-2h1v2zm0-3H5V7h1v2zm0-3H5V5h1v1zm3 9.021H7v-2.04h2v2.04zm-2-3.063v-2h2v2H7zM7 9V7h2v2H7zm2-3.042H7v-1h2v1zm.958-.979H11v1.04H9.958v-1.04zM11 7v2h-1V7h1zm.021 8.021H10v-2.04h1.021v2.04zm0-3.002H10V9.98h1.021v2.039z"/>`),
		g.Group(children),
	)
}