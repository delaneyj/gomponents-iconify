package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.707 3.293L5 2.586l-.707.707L2 5.586l-.707.707l1.414 1.414L3.414 7L4 6.414v11.171L3.414 17l-.707-.707l-1.414 1.414l.707.707l2.293 2.293l.707.707l.707-.707L8 18.414l.707-.707l-1.414-1.414l-.707.707l-.586.585V6.414L6.586 7l.707.707l1.414-1.414L8 5.586L5.707 3.293ZM13 4h-1v2h10V4h-9Zm-2 7h-1v2h12v-2H11Zm1 7h10v2H12v-2Z"/>`),
		g.Group(children),
	)
}