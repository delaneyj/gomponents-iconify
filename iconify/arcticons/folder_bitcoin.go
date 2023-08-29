package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderBitcoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.84 26.5a3.2 3.2 0 0 1 0 6.4h-5.28V20.1h5.28a3.2 3.2 0 0 1 0 6.4h0Zm0 0h-5.28m0 6.4h-1.6m1.6-12.8h-1.6m3.2 14.4v-1.6m3.2 1.6v-1.6m-3.2-12.8v-1.6m3.2 1.6v-1.6"/>`),
		g.Group(children),
	)
}