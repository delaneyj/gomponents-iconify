package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clapboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m2.042 7.954l14.916-3.523l-.98-4.056L1.063 3.898l.979 4.056zM13.55 4.357l-1.448-2.312l2.01-.486l1.449 2.312l-2.011.486zm-3.351.81L8.75 2.855l2.009-.485l1.45 2.312l-2.01.485zm-3.351.81L5.4 3.665l2.01-.485l1.449 2.312l-2.011.485zm-3.35.81L2.049 4.475l2.009-.485l1.449 2.312l-2.009.485zm-1.477 1.25H17v7.95H2.021v-7.95z"/>`),
		g.Group(children),
	)
}