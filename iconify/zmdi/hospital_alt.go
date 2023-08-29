package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 91v42l-43 128l43 128v43H0v-43l43-128L0 133V91h271l31-86l50 19l-24 67h56zM277 283v-43h-64v-64h-42v64h-64v43h64v64h42v-64h64z"/>`),
		g.Group(children),
	)
}