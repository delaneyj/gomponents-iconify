package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Travel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024H128q-53 0-90.5-37.5T0 896V384q0-53 37.5-90.5T128 256h128v-96q0-68 51-114T416 0h192q58 0 109 46t51 114v96h128q53 0 90.5 37.5T1024 384v512q0 53-37.5 90.5T896 1024zM288 448q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm352-256q0-27-18.5-45.5T576 128H448q-26 0-45 18.5T384 192v64h256v-64zm192 448l-256 64l64 192l256-64z"/>`),
		g.Group(children),
	)
}