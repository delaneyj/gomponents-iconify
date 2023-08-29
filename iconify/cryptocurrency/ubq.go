package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ubq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m13.503 15.844l3.318-2.01V5L6.008 11.278v9.103l7.777 4.068l-.282-8.605zm4.994.275l-3.318 2.012V27l10.813-6.321v-9.103l-7.777-4.068l.282 8.611zM16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16z"/><path fill-opacity=".305" fill-rule="nonzero" d="m18.215 7.508l7.777 4.068l-7.493 4.593l-.284-8.661zm-4.43 16.941l-7.777-4.068l7.493-4.594l.284 8.662z"/></g>`),
		g.Group(children),
	)
}