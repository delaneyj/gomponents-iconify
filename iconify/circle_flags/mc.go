package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsMc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMc0)"><path fill="#eee" d="m0 256l258.2-43.3L512 256v256H0z"/><path fill="#d80027" d="M0 0h512v256H0z"/></g>`),
		g.Group(children),
	)
}