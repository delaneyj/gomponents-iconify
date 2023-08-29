package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmOpenEnterpriseLanguages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.586 21.414L27.166 25l-3.582 3.587L25 30l5-5l-5-5zm-3.172 0L16.834 25l3.582 3.587L19 30l-5-5l5-5zM22 6h2v8h-2zm-4 0h2v8h-2zm-4 8h-2c-1.103 0-2-.897-2-2V8c0-1.103.897-2 2-2h2c1.103 0 2 .897 2 2v4c0 1.103-.897 2-2 2zm-2-6v4h2V8h-2zM6 6h2v8H6z"/><path fill="currentColor" d="M10 28H4c-1.103 0-2-.897-2-2V4c0-1.103.897-2 2-2h22c1.103 0 2 .897 2 2v12h-2V4H4v22h6v2z"/>`),
		g.Group(children),
	)
}