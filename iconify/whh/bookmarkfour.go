package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmarkfour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M991 703H703v288q0 22-27 30.5t-40-5.5L26 406Q0 380 0 343t26-63L280 26q26-26 63-26t63 26l610 610q14 13 5.5 40T991 703zM448 224l-89 68l-108-41l41 108l-68 89l115-1l77 96V416h127l-96-77z"/>`),
		g.Group(children),
	)
}