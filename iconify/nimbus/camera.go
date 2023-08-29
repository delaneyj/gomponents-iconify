package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 3.5h-4l-.22-.65A1.24 1.24 0 0 0 9.6 2H6.4a1.24 1.24 0 0 0-1.18.85L5 3.5H1a1 1 0 0 0-1 1V13a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V4.5a1 1 0 0 0-1-1zm-.25 9.25H1.25v-8H5.9l.5-1.5h3.2l.5 1.5h4.65z"/><path fill="currentColor" d="M8 6a2.59 2.59 0 0 0-2.67 2.5A2.59 2.59 0 0 0 8 11a2.59 2.59 0 0 0 2.67-2.5A2.59 2.59 0 0 0 8 6zm0 3.75A1.35 1.35 0 0 1 6.58 8.5A1.35 1.35 0 0 1 8 7.25A1.35 1.35 0 0 1 9.42 8.5A1.35 1.35 0 0 1 8 9.75z"/><ellipse cx="12.84" cy="6.63" fill="currentColor" rx=".67" ry=".63"/>`),
		g.Group(children),
	)
}