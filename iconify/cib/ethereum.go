package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ethereum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m15.927 23.959l-9.823-5.797l9.817 13.839l9.828-13.839l-9.828 5.797zM16.073 0L6.254 16.297l9.819 5.807l9.823-5.801z"/>`),
		g.Group(children),
	)
}