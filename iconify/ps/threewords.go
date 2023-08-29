package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Threewords(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M181 187L303 46V2H22v110h57V54h128l24-2v2q-11 7-18 16L101 203l15 33h10q10-1 17-1q44 0 73 22.5t29 61.5q0 37-27.5 61.5T150 405q-29 0-57-12.5T52 368l-14-13l-34 46q6 7 18 17t51 27t81 17q70 0 114-42.5T312 317t-39.5-92.5T181 187z"/>`),
		g.Group(children),
	)
}