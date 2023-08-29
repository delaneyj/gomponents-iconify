package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 128c70.7 0 128 57.3 128 128s-57.3 128-128 128H192c-70.7 0-128-57.3-128-128s57.3-128 128-128h192zm192 128c0-106-86-192-192-192H192C86 64 0 150 0 256s86 192 192 192h192c106 0 192-86 192-192zm-384 96a96 96 0 1 0 0-192a96 96 0 1 0 0 192z"/>`),
		g.Group(children),
	)
}