package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsTz0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTz0)"><path fill="#eee" d="M0 0h512v512H0z"/><path fill="#ffda44" d="M399 0L167 167L0 399v45l68 68h45l232-167l167-232V68L444 0Z"/><path fill="#333" d="M444 0L0 444v68h68L512 68V0Z"/><path fill="#338af3" d="m113 512l399-399v399z"/><path fill="#6da544" d="M0 399V0h399Z"/></g>`),
		g.Group(children),
	)
}