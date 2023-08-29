package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attachment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5v14H5V3h14v18H9V7h6v10h-2V9h-2v10h6V5H7z"/>`),
		g.Group(children),
	)
}