package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4a4 4 0 0 0-3.999 4.102a1 1 0 0 1-.75.992A3.002 3.002 0 0 0 7 15h1a1 1 0 1 1 0 2H7a5 5 0 0 1-1.97-9.596a6 6 0 0 1 11.169-2.4A6 6 0 0 1 16 17a1 1 0 1 1 0-2a4 4 0 1 0-.328-7.987a1 1 0 0 1-.999-.6A4.001 4.001 0 0 0 11 4zm.293 5.293a1 1 0 0 1 1.414 0l2 2a1 1 0 0 1-1.414 1.414L13 12.414V20a1 1 0 1 1-2 0v-7.586l-.293.293a1 1 0 0 1-1.414-1.414l2-2z"/>`),
		g.Group(children),
	)
}