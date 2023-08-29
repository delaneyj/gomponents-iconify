package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaningGibbousMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32.001 2h.003h-.005C15.431 2 2 15.432 2 32c0 16.568 13.431 30 29.999 30h.005h-.003C48.568 62 62 48.568 62 32C62 15.432 48.568 2 32.001 2zm-.002 58C16.537 60 4 47.463 4 32C4 16.536 16.537 4 31.999 4c2.569 0 5.056.353 7.42 1.001C45.091 9.863 49 20.125 49 32c0 11.875-3.909 22.136-9.581 26.998A28.029 28.029 0 0 1 31.999 60z"/>`),
		g.Group(children),
	)
}