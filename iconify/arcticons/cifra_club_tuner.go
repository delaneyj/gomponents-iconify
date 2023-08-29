package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CifraClubTuner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.729 12v16.7A2.729 2.729 0 0 1 24 31.43h0a2.729 2.729 0 0 1-2.729-2.73V12m-4.096 15.202c-.465-.98-.725-2.075-.725-3.231s.26-2.252.725-3.232m13.65.001c.465.98.725 2.075.725 3.23s-.26 2.252-.725 3.232m-17.442 2.221c-.84-1.634-1.316-3.488-1.316-5.452s.475-3.819 1.316-5.453m21.234 0c.84 1.634 1.316 3.488 1.316 5.453s-.475 3.818-1.316 5.453M24 36v-4.571"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}