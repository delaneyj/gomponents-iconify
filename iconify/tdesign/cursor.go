package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.108 3.114l16.864 5.925l-4.319 3.084l5.707 5.706l-3.536 3.536l-5.706-5.706l-3.085 4.318L3.108 3.114Zm3.269 3.268l3.267 9.3l2.219-3.107l5.96 5.961l.708-.707l-5.961-5.961l3.106-2.219l-9.3-3.267Z"/>`),
		g.Group(children),
	)
}