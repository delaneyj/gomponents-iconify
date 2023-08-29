package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.017 16.708c1.768 1.326 3.241 4.714 2.946 8.397c-.294 2.799-1.767 5.156-2.946 6.04M5.5 17.74v12.816h8.544l11.049 8.839V8.605L14.044 17.74H5.5Zm31.673-6.777c3.094 2.357 5.598 8.397 5.303 15.026c-.442 5.009-2.946 9.133-5.303 10.901"/>`),
		g.Group(children),
	)
}