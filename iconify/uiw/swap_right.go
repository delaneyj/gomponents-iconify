package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.625 12.08H.7c-.176 0-.7.092-.7.71s.524.71.7.71h18.599c.406 0 .701-.287.701-.71a.665.665 0 0 0-.164-.453a4.428 4.428 0 0 0-.173-.187L14.34 6.68c-.348-.278-.668-.27-.96.025c-.292.294-.311.61-.058.95l4.304 4.425Z"/>`),
		g.Group(children),
	)
}