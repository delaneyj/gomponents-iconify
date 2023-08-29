package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adbmanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.3 14.62v14.1m3.52 0h0a3.53 3.53 0 0 1-3.52-3.52v-2.33a3.52 3.52 0 0 1 3.52-3.52h0a3.52 3.52 0 0 1 3.53 3.52v2.33a3.53 3.53 0 0 1-3.53 3.52Zm-7.3-14.1v14.1m-3.52 0h0a3.52 3.52 0 0 1-3.52-3.52v-2.33A3.51 3.51 0 0 1 24 19.35h0a3.51 3.51 0 0 1 3.52 3.52v2.33A3.52 3.52 0 0 1 24 28.72Zm-10.82 0h0a3.53 3.53 0 0 1-3.53-3.52v-2.33a3.52 3.52 0 0 1 3.53-3.52h0a3.52 3.52 0 0 1 3.52 3.52v2.33a3.53 3.53 0 0 1-3.52 3.52Zm3.52-9.37v9.37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}