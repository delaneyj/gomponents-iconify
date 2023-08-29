package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForDenmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm28 30c0 2.375-.299 4.68-.858 6.883H31.017v21.092a27.818 27.818 0 0 1-13.767-4.189V38.883H4.858A27.948 27.948 0 0 1 4 32c0-2.375.299-4.68.858-6.883H17.25V8.215a27.816 27.816 0 0 1 13.767-4.19v21.092h28.125C59.701 27.32 60 29.625 60 32z"/>`),
		g.Group(children),
	)
}