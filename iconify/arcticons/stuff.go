package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stuff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.695 8.5H6.5a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2v-27a2 2 0 0 0-2-2h-.2"/><circle cx="32.5" cy="11.5" r="9.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.626 11.962l2.225 2.418l6.079-5.813m6.57 19.036a19.662 19.662 0 0 1-21.477.343a19.406 19.406 0 0 1-8.815-19.42m4.836 30.968A30.984 30.984 0 0 1 4.5 25.931"/>`),
		g.Group(children),
	)
}