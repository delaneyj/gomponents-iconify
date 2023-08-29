package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelopes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 14H2v-1h13V4h1v10z"/><path fill="currentColor" d="M14 10.77V5.29L9.32 7.47l4.68 3.3zM8.28 7.96L7 8.55l-1.28-.59L0 11.99V12l14-.01l-5.72-4.03zM7 7.45l7-3.27V2H0v2.18l7 3.27zm-2.32.02L0 5.29v5.48l4.68-3.3z"/>`),
		g.Group(children),
	)
}