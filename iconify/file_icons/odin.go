package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Odin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M90.076 452.839L344.858 11.524C704.146 180.817 420.076 682.572 90.076 452.84zM362.953 70.431L149.9 438.64C383.25 563.76 600.708 239.483 362.953 70.43zM301.199 3.273L60.356 423.203C-87.082 228.676 51.866-32.236 301.2 3.273zM224.778 46.5s-89.544 12.35-148.211 98.035C16.934 231.631 61.128 334.43 61.128 334.43L224.778 46.5z"/>`),
		g.Group(children),
	)
}