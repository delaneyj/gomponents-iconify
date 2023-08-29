package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CicsRegionTarget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 7h2v3h-2zm3-3h3v2h-3zm-3-4h2v3h-2zm-4 4h3v2h-3zm-3 12v2.7l-2.5-.7l-.5 1.9l2.4.7L17 23l1.7 1l1.3-2.2l1.3 2.2l1.7-1l-1.4-2.4l2.4-.7l-.5-1.9l-2.5.7V16h-2zm-6 0v-3h3v-2h-3V8h-2v3H8v2h3v3h2z"/><path fill="currentColor" d="M26 26H6V6h13V4H6a2.006 2.006 0 0 0-2 2v20a2.006 2.006 0 0 0 2 2h20a2.006 2.006 0 0 0 2-2V13h-2Z"/>`),
		g.Group(children),
	)
}