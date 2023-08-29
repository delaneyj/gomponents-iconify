package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretCircleDoubleUpThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198.71 57.29A100 100 0 1 0 57.29 198.71A100 100 0 1 0 198.71 57.29Zm-5.66 135.76A92 92 0 1 1 220 128a91.37 91.37 0 0 1-26.95 65.05Zm-30.22-27.88a4 4 0 0 1-5.66 5.66L128 141.66l-29.17 29.17a4 4 0 0 1-5.66-5.66l32-32a4 4 0 0 1 5.66 0Zm0-56a4 4 0 0 1-5.66 5.66L128 85.66l-29.17 29.17a4 4 0 0 1-5.66-5.66l32-32a4 4 0 0 1 5.66 0Z"/>`),
		g.Group(children),
	)
}