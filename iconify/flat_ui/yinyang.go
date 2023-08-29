package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yinyang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="50" cy="50" r="50" fill="#34495E"/><path fill="#fff" d="M50 91C28 91 9 72.645 9 50S28 9 50 9v82z"/><path fill="#34495E" d="M50 9c22 0 41 18.355 41 41S72 91 50 91V9z"/><circle cx="50.5" cy="70.5" r="20.5" fill="#fff"/><circle cx="50.5" cy="70.5" r="8.5" fill="#34495E"/><circle cx="49.5" cy="29.5" r="20.5" fill="#34495E"/><circle cx="50" cy="29.5" r="8.5" fill="#fff"/>`),
		g.Group(children),
	)
}