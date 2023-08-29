package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapboardPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m.979 8.954l14.916-3.523l-.98-4.056L0 4.898l.979 4.056zm11.508-3.597l-1.448-2.312l2.01-.486l1.449 2.312l-2.011.486zm-3.35.81L7.688 3.855l2.009-.485l1.45 2.312l-2.01.485zm-3.352.81L4.337 4.665l2.01-.485l1.449 2.312l-2.011.485zm-3.349.81L.987 5.475l2.009-.485l1.449 2.312l-2.009.485zM1 9.037v8h15v-8H1zm6.904 6.13v-4.25l3.232 2.169l-3.232 2.081z"/>`),
		g.Group(children),
	)
}