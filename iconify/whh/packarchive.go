package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Packarchive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 1024H64q-27 0-45.5-19T0 960V448q0-27 18.5-45.5T64 384h320q3-219 63-353q7-13 19.5-22T489 0h54q11 0 23 9t19 22q57 134 56 353h319q27 0 45.5 18.5T1024 448v512q0 26-19 45t-45 19zm0-576H639q-1 32-4 64h118q15 13 15 32.5T753 578L546 785q-14 15-34.5 15T477 785L270 578q-14-14-14-33.5t14-32.5h117q-2-32-3-64H64v512h896V448z"/>`),
		g.Group(children),
	)
}