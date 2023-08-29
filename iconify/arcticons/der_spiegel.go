package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DerSpiegel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.784 15.61v-8.2a2.909 2.909 0 0 0-2.908-2.91H15.124a2.909 2.909 0 0 0-2.908 2.909V26.21h16.31v10.928h-9.052v-6.687h-7.258v10.14a2.909 2.909 0 0 0 2.908 2.908h17.752a2.909 2.909 0 0 0 2.908-2.909V19.85h-16.31v-8.989h9.052v4.748Z"/>`),
		g.Group(children),
	)
}