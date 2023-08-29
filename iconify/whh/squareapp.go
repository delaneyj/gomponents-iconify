package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squareapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.955 1024h-640q-80 0-136-56.5T.955 832V192q0-80 56-136t136-56h640q79 0 135.5 56t56.5 136v640q0 79-56.5 135.5t-135.5 56.5zm0-800q0-13-9.5-22.5t-22.5-9.5h-576q-13 0-22.5 9.5t-9.5 22.5v576q0 13 9.5 22.5t22.5 9.5h576q13 0 22.5-9.5t9.5-22.5V224zm-224 416h-192q-13 0-22.5-9.5t-9.5-22.5V416q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5z"/>`),
		g.Group(children),
	)
}