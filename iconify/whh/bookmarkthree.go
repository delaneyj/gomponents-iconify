package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmarkthree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M991 703H703v288q0 22-27 30.5t-40-5.5L26 406Q0 380 0 343t26-63L280 26q26-26 63-26t63 26l610 610q14 13 6 40t-31 27z"/>`),
		g.Group(children),
	)
}