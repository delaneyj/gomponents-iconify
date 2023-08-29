package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M86 192H32q-14 0-23-9t-9-23V96q0-14 9-23t23-9h960q14 0 23 9t9 23v64q0 14-9 23t-23 9h-54q-8 201-77.5 332.5T668 698v148q121 42 190.5 171.5T938 1344h54q14 0 23 9t9 23v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h54q10-197 79.5-326.5T356 846V698q-123-42-192.5-173.5T86 192zm784 0H154q3 75 15.5 141.5t39 131.5t75 111T398 639l26 6v254l-26 7q-65 17-113 61.5T210.5 1076T171 1205t-17 139h87l271-271l271 271h87q-4-74-17-139t-39.5-129T739 967.5T626 906l-26-7V645l26-6q66-17 114.5-63t75-111t39-131.5T870 192zM363 502h298L512 650z"/>`),
		g.Group(children),
	)
}