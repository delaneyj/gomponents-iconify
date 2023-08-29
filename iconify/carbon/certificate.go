package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Certificate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24 17l1.912 3.703l4.088.594L27 24l.771 4L24 25.75L20.229 28L21 24l-3-2.703l4.2-.594L24 17zM6 16h6v2H6zm0-4h10v2H6zm0-4h10v2H6z"/><path fill="currentColor" d="M16 26H4V6h24v10h2V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h12Z"/>`),
		g.Group(children),
	)
}