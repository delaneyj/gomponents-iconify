package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neuter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M80 176a112 112 0 1 1 224 0a112 112 0 1 1-224 0zm144 173.1c81.9-15 144-86.8 144-173.1C368 78.8 289.2 0 192 0S16 78.8 16 176c0 86.3 62.1 158.1 144 173.1V480c0 17.7 14.3 32 32 32s32-14.3 32-32V349.1z"/>`),
		g.Group(children),
	)
}