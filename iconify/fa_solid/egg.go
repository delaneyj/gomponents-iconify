package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Egg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M192 0C86 0 0 214 0 320s86 192 192 192s192-86 192-192S298 0 192 0z"/>`),
		g.Group(children),
	)
}