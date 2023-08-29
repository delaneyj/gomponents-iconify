package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Switchon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 768H128q-53 0-90.5-37.5T0 640V128q0-53 37.5-90.5T128 0h768q53 0 90.5 37.5T1024 128v512q0 53-37.5 90.5T896 768zm0-576q0-27-19-45.5T832 128H192q-27 0-45.5 18.5T128 192v384q0 26 18.5 45t45.5 19h640q26 0 45-19t19-45V192zM448 576H256q-27 0-45.5-19T192 512V256q0-27 18.5-45.5T256 192h192q26 0 45 18.5t19 45.5v256q0 26-19 45t-45 19z"/>`),
		g.Group(children),
	)
}