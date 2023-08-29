package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsBd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBd0)"><path fill="#496e2d" d="M0 0h512v512H0z"/><circle cx="200.3" cy="256" r="111.3" fill="#d80027"/></g>`),
		g.Group(children),
	)
}