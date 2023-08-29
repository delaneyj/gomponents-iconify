package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reversocontext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.748 31.263A12.973 12.973 0 0 0 24 11.028m-10.756 5.719A12.973 12.973 0 0 0 24 36.971"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 13.032v4.875l-6.703-6.703L24 4.501h0v8.531zm.001 21.938v-4.875l6.703 6.703l-6.703 6.703V34.97z"/>`),
		g.Group(children),
	)
}