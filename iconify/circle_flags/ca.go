package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ca(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsCa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCa0)"><path fill="#d80027" d="M0 0v512h144l112-64l112 64h144V0H368L256 64L144 0Z"/><path fill="#eee" d="M144 0h224v512H144Z"/><path fill="#d80027" d="m301 289l44-22l-22-11v-22l-45 22l23-44h-23l-22-34l-22 33h-23l23 45l-45-22v22l-22 11l45 22l-12 23h45v33h22v-33h45z"/></g>`),
		g.Group(children),
	)
}