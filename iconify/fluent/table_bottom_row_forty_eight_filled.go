package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableBottomRowFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M42 12.25A6.25 6.25 0 0 0 35.75 6h-23.5A6.25 6.25 0 0 0 6 12.25V29h2.5V12.25a3.75 3.75 0 0 1 3.75-3.75h23.5a3.75 3.75 0 0 1 3.75 3.75V29H42V12.25ZM8.5 40.75a6.24 6.24 0 0 1-2.5-5V31.5h10.5V42h-4.25a6.222 6.222 0 0 1-3.75-1.25ZM42 31.5H31.5V42h4.25a6.222 6.222 0 0 0 3.75-1.25a6.24 6.24 0 0 0 2.5-5V31.5ZM29 42H19V31.5h10V42Z"/>`),
		g.Group(children),
	)
}