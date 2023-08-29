package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newwindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5.896 5.333V21.25h23.417V5.333H5.896zM26.312 18.25H8.896V8.334h17.417v9.916zM4.896 9.542h-3.21V25.46h23.418v-3.21H4.896V9.542z"/>`),
		g.Group(children),
	)
}