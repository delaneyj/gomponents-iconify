package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M35 18a9 9 0 0 0-9-9h-4a9.003 9.003 0 0 0-8.487 6a3 3 0 1 0 5.657 2A3.003 3.003 0 0 1 22 15h4a3 3 0 0 1 2 5.237a2.953 2.953 0 0 0-.085.078l-14 13.528A3 3 0 0 0 16 39h16a3 3 0 1 0 0-6h-8.578l8.626-8.335A8.985 8.985 0 0 0 35 18Zm-9-7a7 7 0 0 1 4.67 12.216L20.252 33.281A1 1 0 0 0 20.948 35H32a1 1 0 1 1 0 2H16a1 1 0 0 1-.695-1.72l14-13.527l-.695-.719l.695.72A4.985 4.985 0 0 0 31 18a5 5 0 0 0-5-5.001h-4a5.002 5.002 0 0 0-4.716 3.333a1 1 0 1 1-1.885-.666A7.002 7.002 0 0 1 22 11h4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}