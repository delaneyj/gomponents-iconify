package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersEightOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 5.5a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5Zm2.857 5.396a5 5 0 1 1-5.714 0a4.25 4.25 0 1 1 5.714 0ZM12 12a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/>`),
		g.Group(children),
	)
}