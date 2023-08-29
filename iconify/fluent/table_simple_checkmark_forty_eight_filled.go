package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleCheckmarkFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 6h10.75v16.75H6V12a6 6 0 0 1 6-6ZM6 25.25h16.75V42H12a6 6 0 0 1-6-6V25.25Zm19.25 0V42H36a6 6 0 0 0 6-6V25.25H25.25Zm0-19.25v16.75H42V12a6 6 0 0 0-6-6H25.25Zm13.13 25.634l-5.5 5.5a1.25 1.25 0 0 1-1.768 0l-2.746-2.746a1.25 1.25 0 0 1 1.768-1.768l1.862 1.862l4.616-4.616a1.25 1.25 0 0 1 1.768 1.768Z"/>`),
		g.Group(children),
	)
}