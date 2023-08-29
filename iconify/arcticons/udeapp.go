package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Udeapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.655 35.892a3.864 3.864 0 0 1 3.712 3.886a3.712 3.712 0 0 1-3.712 3.712M7.44 8.212a3.712 3.712 0 1 1 7.425 0v0l.01 18.79h0a9.038 9.038 0 0 0 9.186 8.884a9.038 9.038 0 0 0 9.185-8.884h0l-.11-18.79a3.712 3.712 0 1 1 7.425 0V26.94a16.56 16.56 0 0 1-33.12 0V8.212"/>`),
		g.Group(children),
	)
}