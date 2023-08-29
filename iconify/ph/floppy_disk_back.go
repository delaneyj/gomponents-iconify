package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDiskBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H91.31A15.86 15.86 0 0 0 80 36.69L36.69 80A15.86 15.86 0 0 0 32 91.31V208a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm0 176H48V91.31L91.31 48H168v32H88a8 8 0 0 0 0 16h80a16 16 0 0 0 16-16V48h24Zm-80-96a40 40 0 1 0 40 40a40 40 0 0 0-40-40Zm0 64a24 24 0 1 1 24-24a24 24 0 0 1-24 24Z"/>`),
		g.Group(children),
	)
}