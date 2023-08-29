package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 7.019v-4a1.008 1.008 0 0 0-.354-.782a.998.998 0 0 0-.829-.22L8.852 4.01A1.002 1.002 0 0 0 8 5.017v9.519A4 4 0 1 0 10 18V9.846L20 8.18v4.355A4 4 0 1 0 22 16V7.019Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}