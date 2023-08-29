package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airbrakedotio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.815.576L24 23.424h-6.072L10.679.576Zm-6.456 0l1.872 5.929l-2.447 7.751c1.038.183 2.09.28 3.144.288c.576 0 1.175-.048 1.824-.096l1.151 3.912a28.7 28.7 0 0 1-2.951.169a26.568 26.568 0 0 1-4.32-.361L5.88 23.424H0L8.181.576Z"/>`),
		g.Group(children),
	)
}