package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkmk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.187 8.738v3.985l4.883-3.157v8.217l1.925 1.111l1.926-1.111V9.57l4.882 3.158V8.742l-6.808-4.269l-6.808 4.265zM12 0l10.375 5.999V18L12 24L1.625 18.006V6.003L12 0z"/>`),
		g.Group(children),
	)
}