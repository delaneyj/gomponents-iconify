package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarksSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M168 72v152l-56-40l-56 40V72a8 8 0 0 1 8-8h96a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M160 56H64a16 16 0 0 0-16 16v152a8 8 0 0 0 12.65 6.51L112 193.83l51.36 36.68A8 8 0 0 0 176 224V72a16 16 0 0 0-16-16Zm0 152.46l-43.36-31a8 8 0 0 0-9.3 0L64 208.45V72h96ZM208 40v152a8 8 0 0 1-16 0V40H88a8 8 0 0 1 0-16h104a16 16 0 0 1 16 16Z"/></g>`),
		g.Group(children),
	)
}