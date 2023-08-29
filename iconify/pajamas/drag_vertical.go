package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.75 7.25a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-2.5Zm-.97 4.85a.75.75 0 0 1 0 1.06l-2.25 2.25l-.53.53l-.53-.53l-2.25-2.25a.75.75 0 1 1 1.06-1.06L8 13.82l1.72-1.72a.75.75 0 0 1 1.06 0ZM6 8a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 6 8ZM1 8a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5A.75.75 0 0 1 1 8Zm9.78-5.28a.75.75 0 1 1-1.06 1.06L8 2.06L6.28 3.78a.75.75 0 0 1-1.06-1.06L7.47.47L8-.06l.53.53l2.25 2.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}