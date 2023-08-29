package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 4l-5 4h10zM8.5 5L5 9.656V27h22V9.656L23.5 5h-5L21 7h1.5L24 9H8l1.5-2H11l2.5-2zM7 11h18v14H7zm5.813 2c-.551.05-.958.543-.907 1.094c.051.55.543.957 1.094.906h6c.36.004.695-.184.879-.496a1.01 1.01 0 0 0 0-1.008c-.184-.312-.52-.5-.879-.496h-6.188z"/>`),
		g.Group(children),
	)
}