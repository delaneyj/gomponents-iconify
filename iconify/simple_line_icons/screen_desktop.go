package simple_line_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenDesktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 95.808H64c-35.184 0-64 28.8-64 64V704c0 35.184 28.816 63.983 64 63.983h416v96.208H320c-17.664 0-32 14.336-32 32s14.336 32 32 32h384c17.664 0 32-14.336 32-32s-14.336-32-32-32H544v-96.208h416c35.184 0 64-28.8 64-63.983V159.808c0-35.2-28.816-64-64-64zM960 704H64V159.808h896V704z"/>`),
		g.Group(children),
	)
}