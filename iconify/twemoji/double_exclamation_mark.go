package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="11" cy="31" r="3" fill="#BE1931"/><path fill="#BE1931" d="M14 23a3 3 0 1 1-6 0V4a3 3 0 1 1 6 0v19z"/><circle cx="25" cy="31" r="3" fill="#BE1931"/><path fill="#BE1931" d="M28 23a3 3 0 0 1-6 0V4a3 3 0 0 1 6 0v19z"/>`),
		g.Group(children),
	)
}