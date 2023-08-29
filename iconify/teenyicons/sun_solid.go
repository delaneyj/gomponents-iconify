package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8 2V0H7v2h1Zm-4.793.498L1.5.792L.793 1.5L2.5 3.206l.707-.708Zm9.293.708L14.207 1.5L13.5.792l-1.707 1.706l.707.708Zm-5 .791a3.499 3.499 0 1 0 0 6.996a3.499 3.499 0 1 0 0-6.996ZM2 6.995H0v1h2v-1Zm13 0h-2v1h2v-1ZM1.5 14.199l1.707-1.707l-.707-.707l-1.707 1.706l.707.708Zm12.707-.708L12.5 11.785l-.707.707L13.5 14.2l.707-.708ZM8 14.99v-1.998H7v1.999h1Z"/>`),
		g.Group(children),
	)
}