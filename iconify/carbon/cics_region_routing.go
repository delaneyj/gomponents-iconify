package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CicsRegionRouting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 2v2h4.6L20 10.6l1.4 1.4L28 5.4V10h2V2h-8zm-3 14v2.7l-2.5-.7l-.5 1.9l2.4.7L17 23l1.7 1l1.3-2.2l1.3 2.2l1.7-1l-1.4-2.4l2.4-.7l-.5-1.9l-2.5.7V16h-2zm-6 0v-3h3v-2h-3V8h-2v3H8v2h3v3h2z"/><path fill="currentColor" d="M26 26H6V6h13V4H6a2.006 2.006 0 0 0-2 2v20a2.006 2.006 0 0 0 2 2h20a2.006 2.006 0 0 0 2-2V13h-2Z"/>`),
		g.Group(children),
	)
}