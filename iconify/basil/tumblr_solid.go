package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TumblrSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.256 17.154a.209.209 0 0 0-.268-.128c-.39.138-.798.215-1.211.229a1.525 1.525 0 0 1-1.676-1.73a.212.212 0 0 0 .002-.03V10.08h3.287a.2.2 0 0 0 .2-.2V7.664a.2.2 0 0 0-.2-.2h-3.287V3.2a.2.2 0 0 0-.2-.2h-2.346a.131.131 0 0 0-.122.13a5.326 5.326 0 0 1-3.306 4.65a.205.205 0 0 0-.129.188V9.88c0 .11.09.2.2.2h1.544v5.702c0 1.944 1.438 4.717 5.23 4.656c1.212 0 2.556-.499 2.965-.944a.17.17 0 0 0 .03-.17l-.713-2.169Z"/>`),
		g.Group(children),
	)
}