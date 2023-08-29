package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Genderless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M144 176c44.1 0 80 35.9 80 80s-35.9 80-80 80s-80-35.9-80-80s35.9-80 80-80m0-64C64.5 112 0 176.5 0 256s64.5 144 144 144s144-64.5 144-144s-64.5-144-144-144z"/>`),
		g.Group(children),
	)
}