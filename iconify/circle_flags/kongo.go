package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kongo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsKongo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKongo0)"><path fill="#eee" d="m0 160l256-32l256 32v192l-256 32L0 352Z"/><path fill="#d80027" d="M0 0h512v160H0Z"/><path fill="#ffda44" d="M0 352h512v160H0Z"/><path fill="#333" d="M144 423L256 79l112 344L75 211h362z"/><path fill="#eee" d="m175 381l81-250l81 250l-213-154h264z"/></g>`),
		g.Group(children),
	)
}