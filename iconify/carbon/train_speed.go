package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainSpeed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 25H2v2h2v2h2v-2h5v2h2v-2h5v2h2v-2h5v2h2v-2h3v-2zm-.286-8.41L18.15 8.64A14.933 14.933 0 0 0 9.652 6H2v2h7.652a12.946 12.946 0 0 1 7.365 2.287l1.036.713H9v2h11.962l7.62 5.238A.966.966 0 0 1 28.033 20H2v2h26.034a2.966 2.966 0 0 0 1.68-5.41z"/>`),
		g.Group(children),
	)
}