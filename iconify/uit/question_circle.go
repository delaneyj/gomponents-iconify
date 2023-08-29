package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.596 8.105A2.498 2.498 0 0 0 9.55 9.897a.5.5 0 1 0 .969.25a1.5 1.5 0 1 1 1.926 1.796a1.507 1.507 0 0 0-.981 1.452v.628a.5.5 0 1 0 1 0v-.628a.517.517 0 0 1 .304-.504a2.498 2.498 0 0 0-.173-4.786zm-.631 7.292a.625.625 0 1 0 0 1.25a.625.625 0 0 0 0-1.25zM12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10zm0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9z"/>`),
		g.Group(children),
	)
}