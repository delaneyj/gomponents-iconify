package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roundrectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768.694 1024h-512q-106 0-181-75t-75-181V256q0-106 75-181t181-75h512q106 0 181 75t75 181v512q0 106-75 181t-181 75zm192-768q0-80-56-136t-136-56h-512q-80 0-136 56t-56 136v512q0 80 56 136t136 56h512q80 0 136-56t56-136V256z"/>`),
		g.Group(children),
	)
}