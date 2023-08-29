package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsJm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsJm0)"><path fill="#333" d="M23.3 488.6L0 465V47.1l23.4-23.7l464 464l24.6-22.6V47.1l-24.5-22.7z"/><path fill="#6da544" d="M23.3 23.3L47.1 0h417.7l23.8 23.4l-464 464L47 512h418l22.6-24.5z"/><path fill="#ffda44" d="M0 0v47.1L208.8 256L0 464.9V512h47.1L256 303.2L464.9 512H512v-47L303.1 256L512 47.2V0h-47.2L256 208.9L47 0z"/></g>`),
		g.Group(children),
	)
}