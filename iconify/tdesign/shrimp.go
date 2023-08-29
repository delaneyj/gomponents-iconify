package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shrimp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.598 5.99L14 9.588V5c1.324 0 2.55.359 3.598.99ZM12 5v6h-2c-2.023 0-3.784-.591-5.02-1.614C3.765 8.38 3 6.913 3 5h9Zm2 9.415l3.59 3.592A7.002 7.002 0 0 1 14 19v-4.585ZM12 17v2H8a2 2 0 0 1 2-2h2Zm2 4c2.19 0 4.215-.798 5.773-2.095A8.939 8.939 0 0 0 23 12a9.027 9.027 0 0 0-3.218-6.898A8.911 8.911 0 0 0 14 3H1v2c0 2.505 1.027 4.538 2.706 5.927C5.366 12.3 7.604 13 10 13h2v2h-2a4 4 0 0 0-4 4v2h8Zm5.153-4.258L15.413 13h5.517a6.916 6.916 0 0 1-1.777 3.742ZM20.928 11h-5.513l3.737-3.735A7.029 7.029 0 0 1 20.928 11ZM9.002 5.998H6.998v2.004h2.004V5.998Z"/>`),
		g.Group(children),
	)
}