package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSweden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm28 30c0 2.375-.3 4.681-.858 6.884H31.017v21.091a27.818 27.818 0 0 1-13.767-4.189V38.884H4.858a28.004 28.004 0 0 1 0-13.767H17.25V8.215a27.795 27.795 0 0 1 13.767-4.189v21.092h28.125C59.7 27.32 60 29.626 60 32z"/>`),
		g.Group(children),
	)
}