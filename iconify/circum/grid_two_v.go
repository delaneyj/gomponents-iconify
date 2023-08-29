package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridTwoV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.436 20.937H15.5a2.5 2.5 0 0 1-2.5-2.5V5.565a2.5 2.5 0 0 1 2.5-2.5h2.933a2.5 2.5 0 0 1 2.5 2.5v12.872a2.5 2.5 0 0 1-2.497 2.5ZM15.5 4.065a1.5 1.5 0 0 0-1.5 1.5v12.872a1.5 1.5 0 0 0 1.5 1.5h2.933a1.5 1.5 0 0 0 1.5-1.5V5.565a1.5 1.5 0 0 0-1.5-1.5Zm-7 16.872H5.564a2.5 2.5 0 0 1-2.5-2.5V5.565a2.5 2.5 0 0 1 2.5-2.5H8.5a2.5 2.5 0 0 1 2.5 2.5v12.872a2.5 2.5 0 0 1-2.5 2.5ZM5.564 4.065a1.5 1.5 0 0 0-1.5 1.5v12.872a1.5 1.5 0 0 0 1.5 1.5H8.5a1.5 1.5 0 0 0 1.5-1.5V5.565a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		g.Group(children),
	)
}