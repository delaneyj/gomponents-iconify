package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsSr0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSr0)"><path fill="#6da544" d="M0 0h512v111.3l-85.3 143.1L512 400.7V512H0V400.7l87-149L0 111.3z"/><path fill="#eee" d="M0 111.3h512V167l-41 84.7l41 93.3v55.7H0V345l44.2-86.6L0 167z"/><path fill="#a2001d" d="M0 167h512v178H0z"/><path fill="#ffda44" d="m256 167l22.1 68h71.5l-57.8 42l22 68l-57.8-42l-57.9 42l22.1-68l-57.8-42h71.5z"/></g>`),
		g.Group(children),
	)
}