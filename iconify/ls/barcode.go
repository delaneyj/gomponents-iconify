package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 644h25V50H0v594zm75 0h49V50H75v594zm98 0h50V50h-50v594zm99 0h100V50H272v594zm149 0h75V50h-75v594zm124 0h50V50h-50v594zm99 0h25V50h-25v594zm75 0h49V50h-49v594z"/>`),
		g.Group(children),
	)
}