package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Urforms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.2 25.4a3.9 3.9 0 0 1 3.9-3.9h0m-3.9-.1v10.5M14.9 16.1v10.5a5.2 5.2 0 1 0 10.4 0V16.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.3 42.5H7.7a2.22 2.22 0 0 1-2.2-2.2V7.7a2.22 2.22 0 0 1 2.2-2.2h32.6a2.22 2.22 0 0 1 2.2 2.2v32.6a2.22 2.22 0 0 1-2.2 2.2Z"/>`),
		g.Group(children),
	)
}