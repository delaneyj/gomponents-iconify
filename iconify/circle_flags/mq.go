package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsMq0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMq0)"><path fill="#6da544" d="M0 0h512v256l-256 44Z"/><path fill="#333" d="M210 256h302v256H0z"/><path fill="#d80027" d="M0 0v512l256-256L0 0z"/></g>`),
		g.Group(children),
	)
}