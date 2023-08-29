package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Industry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M448 0q26 0 45 19t19 45v891l536-429q17-14 40-14q26 0 45 19t19 45v379l536-429q17-14 40-14q26 0 45 19t19 45v1152q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h384z"/>`),
		g.Group(children),
	)
}