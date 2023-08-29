package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M416 1120v224q0 26-19 45t-45 19H96q-26 0-45-19t-19-45v-224q0-26 19-45t45-19h256q26 0 45 19t19 45zM446 64l-28 768q-1 26-20.5 45T352 896H96q-26 0-45.5-19T30 832L2 64Q1 38 19.5 19T64 0h320q26 0 44.5 19T446 64z"/>`),
		g.Group(children),
	)
}