package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icecreamalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 704H64q-27 0-45.5-19T0 640V256Q0 150 75 75T256 0t181 75t75 181v384q0 27-19 45.5T448 704zM320 960q0 26-19 45t-45.5 19t-45-19t-18.5-45V768h128v192z"/>`),
		g.Group(children),
	)
}