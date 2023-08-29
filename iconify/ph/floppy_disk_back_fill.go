package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDiskBackFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32h-20a4 4 0 0 0-4 4v44a16 16 0 0 1-16 16H88.27A8.17 8.17 0 0 1 80 88.53A8 8 0 0 1 88 80h76a4 4 0 0 0 4-4V36a4 4 0 0 0-4-4H91.31A15.86 15.86 0 0 0 80 36.69L36.69 80A15.86 15.86 0 0 0 32 91.31V208a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-80 152a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}