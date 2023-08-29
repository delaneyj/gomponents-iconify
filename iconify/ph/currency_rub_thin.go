package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRubThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M148 36H88a4 4 0 0 0-4 4v100H56a4 4 0 0 0 0 8h28v24H56a4 4 0 0 0 0 8h28v36a4 4 0 0 0 8 0v-36h52a4 4 0 0 0 0-8H92v-24h56a56 56 0 0 0 0-112Zm0 104H92V44h56a48 48 0 0 1 0 96Z"/>`),
		g.Group(children),
	)
}