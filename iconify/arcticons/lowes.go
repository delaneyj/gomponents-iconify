package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lowes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.569 21.808v-2.594h-6.33L24 14.722l-10.239 4.492H7.446v2.594H4.5v11.47h39v-11.47h-2.931z"/>`),
		g.Group(children),
	)
}