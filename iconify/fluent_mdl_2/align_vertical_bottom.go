package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1792v128H0v-128h2048zm-384-128h-512V512h512v1152zM1536 640h-256v896h256V640zM896 1664H384V0h512v1664zM768 128H512v1408h256V128z"/>`),
		g.Group(children),
	)
}