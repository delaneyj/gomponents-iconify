package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 3.5a2.8 2.8 0 0 1 2 .82l5.75 5.75a2.78 2.78 0 0 1 0 3.95L26 19.77a2.76 2.76 0 0 1-3.94 0L16.28 14a2.78 2.78 0 0 1 0-3.95L22 4.32a2.8 2.8 0 0 1 2-.82Zm-11.95 12a2.8 2.8 0 0 1 2 .82L19.77 22a2.78 2.78 0 0 1 0 3.95L14 31.72a2.78 2.78 0 0 1-3.95 0L4.32 26a2.79 2.79 0 0 1 0-3.95l5.75-5.75a2.82 2.82 0 0 1 1.98-.85Zm23.9 0a2.82 2.82 0 0 1 2 .82L43.68 22a2.79 2.79 0 0 1 0 3.95l-5.75 5.75a2.78 2.78 0 0 1-4 0l-5.7-5.7a2.78 2.78 0 0 1 0-3.95L34 16.27a2.8 2.8 0 0 1 2-.82Zm-12 12a2.75 2.75 0 0 1 2 .82L31.72 34a2.78 2.78 0 0 1 0 4L26 43.68a2.78 2.78 0 0 1-3.94 0l-5.75-5.75a2.78 2.78 0 0 1 0-4l5.69-5.7a2.75 2.75 0 0 1 2-.82Z"/>`),
		g.Group(children),
	)
}