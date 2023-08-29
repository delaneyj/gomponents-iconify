package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3v-1.586l2.44 2.44c.944.945 2.56.275 2.56-1.061V8.207c0-1.336-1.616-2.006-2.56-1.06L17 9.585V8a3 3 0 0 0-3-3H5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}