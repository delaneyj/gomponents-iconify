package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaClipboardFill0"><g id="evaClipboardFill1"><g id="evaClipboardFill2" fill="currentColor"><path d="M18 4v3a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V4a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Z"/><rect width="10" height="6" x="7" y="2" rx="1" ry="1"/></g></g></g>`),
		g.Group(children),
	)
}