package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Floobits(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M436.363 327.866c-33.516-19.322-33.516-67.89 0-87.211C469.877 221.333 512 245.616 512 284.26s-42.122 62.928-75.637 43.605zM25.137 267.018c-33.516-19.322-33.516-67.89 0-87.212c33.515-19.322 75.637 4.962 75.637 43.606s-42.122 62.928-75.637 43.606zm196.731-119.326h97.61L441.55 425.088h-97.92zM75.703 86.912h97.587L295.208 364.14h-97.84z"/>`),
		g.Group(children),
	)
}