package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NaturePeople(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M430 156q0 57-37.5 99T299 304v83h64v42H21V323H0v-86q0-8 6.5-14.5T21 216h64q9 0 15.5 6.5T107 237v86H85v64h171v-84q-53-9-88.5-50.5T132 156q0-62 43.5-106T281 6t105.5 44T430 156zM53.5 195q-13.5 0-23-9.5t-9.5-23t9.5-22.5t23-9t22.5 9t9 22.5t-9 23t-22.5 9.5z"/>`),
		g.Group(children),
	)
}