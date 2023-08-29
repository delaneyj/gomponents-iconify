package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuoiTre(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.07 25.06L19.61 5.5v10.031h3.01v9.53h-3.01v7.022c.501.867 2.507.87 3.01 0v6.52c0 5.02-11.034 6.022-11.034-2.006V25.06H6.068Zm19.31 0L38.92 5.5v10.031h3.01v9.53h-3.01v7.022c.501.867 2.507.87 3.01 0v6.52c0 5.02-11.035 6.022-11.035-2.006V25.06H25.38Z"/>`),
		g.Group(children),
	)
}