package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFrenchGuiana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffce31" d="M2 32c0 16.6 13.4 30 30 30c11.7 0 21.9-6.8 26.8-16.6L5.2 18.6C3.1 22.6 2 27.2 2 32z"/><path fill="#699635" d="M32 2C20.3 2 10.1 8.8 5.2 18.6l53.7 26.8c2-4 3.2-8.6 3.2-13.4C62 15.4 48.6 2 32 2z"/><path fill="#da121a" d="M32 17.9L35.3 28h10.6l-8.6 6.3l3.3 10.1l-8.6-6.2l-8.6 6.2l3.3-10.1l-8.6-6.3h10.6z"/>`),
		g.Group(children),
	)
}