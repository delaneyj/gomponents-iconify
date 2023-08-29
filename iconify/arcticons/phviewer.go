package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phviewer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.3 42.5H7.7a2.22 2.22 0 0 1-2.2-2.2V7.7a2.22 2.22 0 0 1 2.2-2.2h32.6a2.22 2.22 0 0 1 2.2 2.2v32.6a2.22 2.22 0 0 1-2.2 2.2Zm-14.08-30v17.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.22 22.6a4.31 4.31 0 0 1 4.3-4.3h0a4.31 4.31 0 0 1 4.3 4.3v7.1m-21.64-4.3a4.31 4.31 0 0 0 4.3 4.3h0a4.31 4.31 0 0 0 4.3-4.3v-2.8a4.31 4.31 0 0 0-4.3-4.3h0a4.31 4.31 0 0 0-4.3 4.3m0-4.3v17.2"/>`),
		g.Group(children),
	)
}