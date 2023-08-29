package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.78 3.78a.75.75 0 0 0 0-1.06L8.53.47L8-.06l-.53.53l-2.25 2.25a.75.75 0 0 0 1.06 1.06L8 2.06l1.72 1.72a.75.75 0 0 0 1.06 0ZM3.84 5.22a.75.75 0 0 0-1.06 0L.53 7.47L0 8l.53.53l2.25 2.25a.75.75 0 1 0 1.06-1.06L2.12 8l1.72-1.72a.75.75 0 0 0 0-1.06Zm6.94 6.88a.75.75 0 0 1 0 1.06l-2.25 2.25l-.53.53l-.53-.53l-2.25-2.25a.75.75 0 1 1 1.06-1.06L8 13.82l1.72-1.72a.75.75 0 0 1 1.06 0Zm2.44-6.88a.75.75 0 0 0-1.06 1.06L13.88 8l-1.72 1.72a.75.75 0 1 0 1.06 1.06l2.25-2.25L16 8l-.53-.53l-2.25-2.25ZM8 6a.75.75 0 0 1 .75.75v.5h.5a.75.75 0 0 1 0 1.5h-.5v.5a.75.75 0 1 1-1.5 0v-.5h-.5a.75.75 0 1 1 0-1.5h.5v-.5A.75.75 0 0 1 8 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}