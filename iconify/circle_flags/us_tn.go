package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsTn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsUsTn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsTn0)"><path fill="#0052b4" d="M448 0h64v512h-64l-16-256Z"/><path fill="#eee" d="M416 0h32v512h-32l-16-256Z"/><path fill="#d80027" d="M0 0h416v512H0z"/><circle cx="208" cy="256" r="160" fill="#eee"/><circle cx="208" cy="256" r="128" fill="#0052b4"/><path fill="#eee" d="m145 284l60 83V265l-60 83l98-32zm15-128l-22 100l76-68l-102 11l89 51zm147 49l-42 93l-21-100l75 69l-102-11z"/></g>`),
		g.Group(children),
	)
}