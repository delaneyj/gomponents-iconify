package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCardMiniFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2h12a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-8.58a1 1 0 0 1 .292-.707l1.562-1.567A.5.5 0 0 0 6 9.793V3a1 1 0 0 1 1-1Zm8 2v4h2V4h-2Zm-3 0v4h2V4h-2ZM9 4v4h2V4H9Z"/>`),
		g.Group(children),
	)
}