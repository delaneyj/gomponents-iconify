package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transnistria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsTransnistria0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsTransnistria0)"><path fill="#a2001d" d="M0 0h512v189.5l-39 62l39 71.6V512H0V323l40.8-67L0 189.5z"/><path fill="#6da544" d="M0 189.5h512v133.6H0z"/></g>`),
		g.Group(children),
	)
}