package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.672 1.62A5.534 5.534 0 0 1 7.585 0h.207C10.122 0 12 1.925 12 4.243a4.15 4.15 0 0 1-2.276 3.704A3.118 3.118 0 0 0 8 10.737V12H7v-1.264c0-1.56.881-2.986 2.276-3.683A3.15 3.15 0 0 0 11 4.243C11 2.465 9.558 1 7.792 1h-.207a4.534 4.534 0 0 0-3.206 1.328l-.525.526l-.708-.708l.526-.525ZM8 15H7v-1h1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}