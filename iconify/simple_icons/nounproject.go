package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nounproject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.672 8.846H24v6.327h-6.328zM6.328 11.99a3.164 3.164 0 0 1-3.164 3.163A3.164 3.164 0 0 1 0 11.991a3.164 3.164 0 0 1 3.164-3.164a3.164 3.164 0 0 1 3.164 3.164m5.504 1.142l2.04 2.021l1.142-1.16l-2.022-2.003l2.022-2.003l-1.142-1.142l-2.04 2.003L9.81 8.846L8.649 9.988l2.022 2.003l-2.022 2.003l1.16 1.16Z"/>`),
		g.Group(children),
	)
}