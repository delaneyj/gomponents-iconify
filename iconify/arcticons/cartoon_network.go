package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CartoonNetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.132 30.753V17.247L42.5 30.753V17.247m-22.158 8.977v.055a4.474 4.474 0 0 1-4.474 4.474h0a4.474 4.474 0 0 1-4.474-4.474v-4.558a4.474 4.474 0 0 1 4.474-4.474h0a4.474 4.474 0 0 1 4.474 4.474v.055"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 13.632h20.736v20.736H5.5z"/>`),
		g.Group(children),
	)
}