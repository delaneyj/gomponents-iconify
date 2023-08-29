package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GbSct(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsGbSct0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGbSct0)"><path fill="#0052b4" d="M0 68L68 0h376l68 68v376l-68 68H68L0 444Z"/><path fill="#eee" d="M0 0v68l188 188L0 444v68h68l188-188l188 188h68v-68L324 256L512 68V0h-68L256 188L68 0H0z"/></g>`),
		g.Group(children),
	)
}