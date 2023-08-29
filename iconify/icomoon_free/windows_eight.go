package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowsEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M.005 8L0 3.124l6-.815V8zM7 2.164L14.998 1v7H7zM15 9l-.002 7L7 14.875V9zm-9 5.747l-5.995-.822V8.999H6z"/>`),
		g.Group(children),
	)
}