package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<ellipse cx="32" cy="31.999" fill="#e0e7ec" rx="25.02" ry="25.01"/><path fill="#405866" d="M32 1C14.91 1 1 14.904 1 32c0 17.09 13.904 30.998 31 30.998c17.1 0 31-13.904 31-30.998C63 14.91 49.096 1 32 1m0 54.48c-12.947 0-23.485-10.534-23.485-23.481C8.515 19.052 19.05 8.51 32 8.51c12.947 0 23.486 10.54 23.486 23.49c0 12.947-10.534 23.48-23.486 23.48"/>`),
		g.Group(children),
	)
}