package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsVc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsVc0)"><path fill="#ffda44" d="M378.5 0H133.7l-22.3 256l22.3 256h244.8l22.3-256z"/><path fill="#338af3" d="M133.7 512V0H0v512z"/><path fill="#6da544" d="M512 0H378.5v512H512zM200.4 322.8L156 256l44.5-66.8l44.7 66.8zm111.4 0L267.1 256l44.6-66.8l44.5 66.8zm-55.7 89L211.6 345l44.5-66.7l44.5 66.7z"/></g>`),
		g.Group(children),
	)
}