package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyMagenta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 28.696v-9.403l4.707 9.414l4.706-9.399v9.399m2.174-3.883a2.353 2.353 0 0 1 2.353-2.354h0a2.353 2.353 0 0 1 2.353 2.353v3.883m-4.706-6.236v6.236m4.707-3.882a2.353 2.353 0 0 1 2.353-2.354h0a2.353 2.353 0 0 1 2.353 2.354v3.882"/>`),
		g.Group(children),
	)
}