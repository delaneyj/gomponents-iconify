package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsBs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBs0)"><path fill="#338af3" d="M0 0h512v167l-37.4 89l37.4 89v167H0l49.6-252z"/><path fill="#ffda44" d="M108.3 167H512v178H108.3z"/><path fill="#333" d="M0 0v512l256-256L0 0z"/></g>`),
		g.Group(children),
	)
}