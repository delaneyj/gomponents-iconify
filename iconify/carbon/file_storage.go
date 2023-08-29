package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileStorage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 20h-2v2h2v6H4v-6h2v-2H4a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2Z"/><circle cx="7" cy="25" r="1" fill="currentColor"/><path fill="currentColor" d="m22.707 7.293l-5-5A1 1 0 0 0 17 2h-6a2.002 2.002 0 0 0-2 2v16a2.002 2.002 0 0 0 2 2h10a2.002 2.002 0 0 0 2-2V8a1 1 0 0 0-.293-.707ZM20.586 8H17V4.414ZM11 20V4h4v4a2.002 2.002 0 0 0 2 2h4v10Z"/>`),
		g.Group(children),
	)
}