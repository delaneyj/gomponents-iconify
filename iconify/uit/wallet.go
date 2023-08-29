package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 7H18V6a3.003 3.003 0 0 0-3-3H4.5A2.502 2.502 0 0 0 2 5.5V18a3.003 3.003 0 0 0 3 3h14.5a2.502 2.502 0 0 0 2.5-2.5v-9A2.502 2.502 0 0 0 19.5 7zm-15-3H15a2.003 2.003 0 0 1 2 2v1H4.5a1.5 1.5 0 1 1 0-3zM21 16h-2a2 2 0 0 1 0-4h2v4zm0-5h-2a3 3 0 1 0 0 6h2v1.5a1.5 1.5 0 0 1-1.5 1.5H5a2.003 2.003 0 0 1-2-2V7.499c.432.326.959.502 1.5.501h15A1.5 1.5 0 0 1 21 9.5V11z"/>`),
		g.Group(children),
	)
}