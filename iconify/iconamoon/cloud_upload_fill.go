package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUploadFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.697 6.697a7.5 7.5 0 0 1 12.794 4.927A4.002 4.002 0 0 1 18.5 19.5h-12a5 5 0 0 1-1.667-9.715a7.47 7.47 0 0 1 1.864-3.088Zm6.01 1.596a1 1 0 0 0-1.414 0l-2 2a1 1 0 1 0 1.414 1.414l.293-.293V15a1 1 0 1 0 2 0v-3.586l.293.293a1 1 0 0 0 1.414-1.414l-2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}