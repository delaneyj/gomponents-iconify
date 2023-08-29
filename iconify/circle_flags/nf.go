package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsNf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNf0)"><path fill="#6da544" d="M0 0h144.7l108.4 41.2L367.3 0H512v512H367.3l-110.2-41.4L144.7 512H0z"/><path fill="#eee" d="M144.7 0h222.6v512H144.7z"/><path fill="#6da544" d="m323 334l-67-211.6L189.3 334h50.1v55.7h33.4V334z"/></g>`),
		g.Group(children),
	)
}