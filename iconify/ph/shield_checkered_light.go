package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldCheckeredLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 42H48a14 14 0 0 0-14 14v58.77c0 88.25 74.68 117.52 89.65 122.49a13.5 13.5 0 0 0 8.7 0c15-5 89.65-34.24 89.65-122.49V56a14 14 0 0 0-14-14Zm2 14v58.77c0 2.46-.07 4.86-.19 7.21H134V54h74a2 2 0 0 1 2 2ZM46 56a2 2 0 0 1 2-2h74v68H46.19c-.12-2.35-.19-4.75-.19-7.21Zm1.4 78H122v89.9c-18.83-7.46-65.92-31.72-74.6-89.9Zm86.6 89.9V134h74.6c-8.68 58.17-55.77 82.44-74.6 89.9Z"/>`),
		g.Group(children),
	)
}