package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M160 192q-13 0-22.5-9.5T128 160t9.5-22.5T160 128h704q13 0 22.5 9.5T896 160t-9.5 22.5T864 192H160zm672 832H160q-66 0-113-47T0 864V160Q0 94 47 47T160 0h704q13 0 22.5 9.5T896 32t-9.5 22.5T864 64H160q-40 0-68 28t-28 68t28 68t68 28h736v704q0 27-18.5 45.5T832 1024z"/>`),
		g.Group(children),
	)
}