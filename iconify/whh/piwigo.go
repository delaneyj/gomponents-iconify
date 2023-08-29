package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Piwigo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 704H64q-26 0-45-19T0 640V128q0-27 19-45.5T64 64h384q0-27 18.5-45.5T512 0h320q26 0 45 18.5T896 64h64q26 0 45 18.5t19 45.5v512q0 27-18.5 45.5T960 704zM320 192q-80 0-136 56t-56 135.5t56 136T320 576t136-56.5t56-136T456 248t-136-56zM768 64H576q-13 0-22.5 9.5T544 96t9.5 22.5T576 128h192q13 0 22.5-9.5T800 96t-9.5-22.5T768 64zM320 512q-53 0-90.5-37.5T192 384t37.5-90.5T320 256t90.5 37.5T448 384t-37.5 90.5T320 512z"/>`),
		g.Group(children),
	)
}