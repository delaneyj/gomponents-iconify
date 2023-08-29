package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autodoc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.894 5.5v8.99a3.021 3.021 0 0 0 3.02 3.02H42.5v11.022H22.234a3.021 3.021 0 0 1-3.021-3.021V5.5h10.68Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.787 28.532V42.5h-10.68v-8.99a3.021 3.021 0 0 0-3.022-3.02H5.5V19.467h13.713"/>`),
		g.Group(children),
	)
}