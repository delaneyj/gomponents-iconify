package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertTextBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 388v1524H268v-384h116v268h1548V504h-780V388h896zM374 1272H0v-44q19-3 38-5t37-8q14-4 30-26t32-49t27-53t19-41l388-918h82l382 928q9 22 21 50t29 53t38 44t49 22q28 3 56 3v44H782v-44h37q20 0 37-3t30-16t12-37q0-22-11-60t-26-80t-30-79t-26-62H359q-9 23-26 61t-34 80t-29 79t-13 60q0 24 12 35t30 16t38 4t37 2v44zm28-509h344L570 385L402 763z"/>`),
		g.Group(children),
	)
}