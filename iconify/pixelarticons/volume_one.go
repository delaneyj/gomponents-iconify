package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 2h-2v2h-2v2H9v2H5v8h4v2h2v2h2v2h2V2zm-4 16v-2H9v-2H7v-4h2V8h2V6h2v12h-2zm6-8h2v4h-2v-4z"/>`),
		g.Group(children),
	)
}