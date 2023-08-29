package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageTextAnalysis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 22h6v2h-6z"/><path fill="currentColor" d="M2 18v2h2v8c0 1.103.897 2 2 2h20c1.103 0 2-.897 2-2v-8h2v-2H2zm24 10H6v-8h20v8zm3-12h-5c-1.103 0-2-.897-2-2V8c0-1.103.897-2 2-2h5v2h-5v6h5v2zM18 6h-4V2h-2v14h6c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2zm-4 8V8h4v6h-4zM8 6H3v2h5v2H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h6V8c0-1.103-.897-2-2-2zm0 8H4v-2h4v2z"/>`),
		g.Group(children),
	)
}