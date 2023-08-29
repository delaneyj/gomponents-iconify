package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024H256q-106 0-181-75T0 768q0-26 19-45t45.5-19t45 19t18.5 45q0 53 37.5 90.5T256 896h256q53 0 91-37.5t38-90.5V605q-60 35-129 35H256q-106 0-181-75T0 384V256Q0 150 75 75T256 0h256q107 0 182 75t75 181v512q0 106-75 181t-182 75zm129-768q0-53-37.5-90.5T512 128H256q-53 0-90.5 37.5T128 256v128q0 53 37.5 90.5T256 512h256q53 0 91-37.5t38-90.5V256z"/>`),
		g.Group(children),
	)
}