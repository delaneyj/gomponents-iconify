package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6.929.757L4.383 5h6.234L8.07.757l.86-.514L11.783 5h2.514c.388 0 .703.315.703.703v.439a18.96 18.96 0 0 1-2.002 8.48a.684.684 0 0 1-.612.378H2.614a.685.685 0 0 1-.612-.379A18.96 18.96 0 0 1 0 6.141v-.438C0 5.315.315 5 .703 5h2.514L6.07.243l.858.514Z"/>`),
		g.Group(children),
	)
}