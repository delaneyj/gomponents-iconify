package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torigate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 384v128H923q37 236 37 512H832q0-269-44-512H236q-44 244-44 512H64q0-276 37-512H0V384h124q17-75 32-128H64L0 0h1024l-64 256h-92q15 52 32 128h124zm-760 0h184V256H301q-18 51-37 128zm459-128H576v128h184q-19-76-37-128z"/>`),
		g.Group(children),
	)
}