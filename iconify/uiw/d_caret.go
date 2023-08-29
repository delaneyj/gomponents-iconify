package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DCaret(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.407 11.333c.618 0 .773.64.37 1.054l-5.23 5.39a.786.786 0 0 1-1.096 0l-5.229-5.39c-.402-.414-.246-1.054.37-1.054Zm-4.86-9.11l5.23 5.39c.403.414.248 1.054-.37 1.054H4.592c-.616 0-.772-.64-.37-1.054l5.229-5.39a.786.786 0 0 1 1.097 0Z"/>`),
		g.Group(children),
	)
}