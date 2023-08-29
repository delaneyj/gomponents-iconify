package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AqTrueSouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsAqTrueSouth0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAqTrueSouth0)"><path fill="#eee" d="M0 256L256 28l256 228v256H0Z"/><path fill="#0052b4" d="m114 256l142 228l142-228l-142 57ZM0 0h512v256H398L256 28L114 256H0Z"/></g>`),
		g.Group(children),
	)
}