package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CxFileExplorer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.366 19.984L34 33.016m0-13.032l-8.634 13.032m-2.732-4.37v.053a4.317 4.317 0 0 1-4.317 4.317h0A4.317 4.317 0 0 1 14 28.7v-4.398a4.317 4.317 0 0 1 4.317-4.317h0a4.317 4.317 0 0 1 4.317 4.317v.053"/>`),
		g.Group(children),
	)
}