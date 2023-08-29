package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Queue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M23 5v8m-4-4h8M7 9.731V24.27c0 1.508 1.753 2.293 2.83 1.266l7.632-7.269a1.758 1.758 0 0 0 0-2.532L9.83 8.465C8.753 7.44 7 8.223 7 9.731Z"/>`),
		g.Group(children),
	)
}