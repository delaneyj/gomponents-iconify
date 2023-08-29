package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weechat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.457 14.898a21.593 21.593 0 0 0 1.831-12.186a21.594 21.594 0 0 0-12.186 1.83a21.3 21.3 0 0 0-18.204 0a21.595 21.595 0 0 0-12.186-1.83a21.593 21.593 0 0 0 1.83 12.186A21.492 21.492 0 1 0 45.5 24a21.378 21.378 0 0 0-2.043-9.102Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14 30.145a3.355 3.355 0 1 1 6.71 0Zm13.29 0a3.355 3.355 0 0 1 6.71 0Z"/>`),
		g.Group(children),
	)
}