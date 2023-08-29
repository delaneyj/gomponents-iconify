package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExportRtlTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.246 4.504a.75.75 0 0 0-.743.648l-.007.102v13.498a.75.75 0 0 0 1.493.102l.007-.101v-13.5a.75.75 0 0 0-.75-.75ZM8.786 6.387l-.083-.094a1 1 0 0 0-1.32-.083l-.094.083l-4.997 4.998a1 1 0 0 0-.083 1.32l.083.093l4.997 5.004a1 1 0 0 0 1.498-1.32l-.083-.094L5.415 13h12.581a1 1 0 0 0 .993-.883l.007-.117a1 1 0 0 0-.883-.993L17.996 11H5.412l3.291-3.293a1 1 0 0 0 .083-1.32l-.083-.094l.083.094Z"/>`),
		g.Group(children),
	)
}