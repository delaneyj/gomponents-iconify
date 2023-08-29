package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoEntry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#fff" d="M15.58 51.47h96.83v25.05H15.58z"/><path fill="#ed6c30" d="M64 .57C28.96.57.57 28.97.57 64c0 35.04 28.39 63.43 63.43 63.43c35.02 0 63.43-28.4 63.43-63.43S99.02.57 64 .57zM15.58 76.53V51.47h96.83v25.05H15.58z"/>`),
		g.Group(children),
	)
}