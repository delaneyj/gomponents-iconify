package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M69 5Q42 5 23.5 23.5T5 69v342q0 27 18.5 45.5T69 475h342q27 0 45.5-18.5T475 411V170L357 5H69zm150 43v64h42V48h22v85H133V48h86zm-86 384V304h214v128H133zm299-21q0 9-6 15t-15 6h-22V304q0-17-12.5-30T347 261H133q-17 0-29.5 13T91 304v128H69q-9 0-15-6t-6-15V69q0-9 6-15t15-6h22v85q0 18 12.5 30.5T133 176h150q17 0 29.5-12.5T325 133V48h11l96 134v229z"/>`),
		g.Group(children),
	)
}