package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromeRestore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M3 5v9h9V5H3zm8 8H4V6h7v7z"/><path fill-rule="evenodd" d="M5 5h1V4h7v7h-1v1h2V3H5v2z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}