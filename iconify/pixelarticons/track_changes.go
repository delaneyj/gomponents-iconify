package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrackChanges(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2H2v20h20V4h-2v16H4V4h7v2H6v12h12V8h-2v8H8V8h3v2h-1v4h4v-4h-1V2h-2z"/>`),
		g.Group(children),
	)
}