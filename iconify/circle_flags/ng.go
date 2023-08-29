package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ng(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsNg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNg0)"><path fill="#6da544" d="M0 0v512h160l96-64l96 64h160V0H352l-96 64l-96-64Z"/><path fill="#eee" d="M160 0h192v512H160Z"/></g>`),
		g.Group(children),
	)
}