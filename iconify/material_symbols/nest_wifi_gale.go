package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestWifiGale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.6 10h18.8l-.275-4.125q-.05-.8-.625-1.338T19.125 4H4.875q-.8 0-1.375.537t-.625 1.338L2.6 10Zm1.55 9h.525L5 20h14l.325-1h.525q.875 0 1.462-.625t.538-1.5L21.525 12H2.475l-.325 4.875q-.05.875.537 1.5T4.15 19Z"/>`),
		g.Group(children),
	)
}