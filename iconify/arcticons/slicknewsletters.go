package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slicknewsletters(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.755 39.535L6.31 26.307c-2.236-1.703-2.498-6.703-.239-8.915L19.87 8.464L43.5 23.755l-12.067 7.808a4.779 4.779 0 0 0-2.41 4.207c-.008 1.467.848 3.439 2.41 3.038L43.5 30.999l-5.325-3.445"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.87 8.464l1.351 13.43l22.279 1.86"/>`),
		g.Group(children),
	)
}