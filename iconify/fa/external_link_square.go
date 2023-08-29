package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExternalLinkSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1280 800V320q0-26-19-45t-45-19H736q-42 0-59 39q-17 41 14 70l144 144l-534 534q-19 19-19 45t19 45l102 102q19 19 45 19t45-19l534-534l144 144q18 19 45 19q12 0 25-5q39-17 39-59zm256-512v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288z"/>`),
		g.Group(children),
	)
}