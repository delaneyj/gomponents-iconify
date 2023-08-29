package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeafBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.25 2.083a3.46 3.46 0 0 0-.713.243C6.553 4.188 4 9.395 4 13.856c0 4.24 3.183 7.724 7.25 8.109V2.083Zm1.5 19.882c4.067-.385 7.25-3.868 7.25-8.108c0-.406-.021-.818-.063-1.234l-7.187 7.188v2.154ZM18.26 7.18a13.427 13.427 0 0 0-1.34-2.04l-4.17 4.17v3.38l5.51-5.51Zm-2.352-3.15a9.15 9.15 0 0 0-2.445-1.704a3.46 3.46 0 0 0-.713-.243v5.106l3.158-3.158Zm3.028 4.594l-6.186 6.187v2.878l6.75-6.75l.132-.132a15.021 15.021 0 0 0-.696-2.183Z"/>`),
		g.Group(children),
	)
}