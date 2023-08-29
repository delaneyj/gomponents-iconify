package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polygon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M28 10c0 .507-.094.992-.266 1.438l7.696 5.497a4 4 0 1 1 2.706 7.063l-2.32 10.437A4 4 0 1 1 30.126 39H17.874A4.002 4.002 0 0 1 10 38a4 4 0 0 1 2.184-3.565l-2.32-10.437a4 4 0 1 1 2.706-7.063l7.696-5.497A4 4 0 1 1 28 10Zm-6.57 3.065l-7.696 5.497a4 4 0 0 1-1.917 5.003l2.319 10.437A4.002 4.002 0 0 1 17.874 37h12.252a4.002 4.002 0 0 1 3.738-2.998l2.32-10.437a4 4 0 0 1-1.918-5.003l-7.695-5.497A3.984 3.984 0 0 1 24 14a3.984 3.984 0 0 1-2.57-.935Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}