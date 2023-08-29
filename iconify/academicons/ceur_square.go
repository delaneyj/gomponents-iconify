package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CeurSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M48 32C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48H48zm16 64h106.668v53.334h-53.334v213.332H224V416H64V96zm160 0h160v320H277.332v-53.334h53.334V149.334H224V96z"/>`),
		g.Group(children),
	)
}