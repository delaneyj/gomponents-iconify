package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ubereats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Zm-29-7.73h4.68M11.5 25.4h4.68m-4.68 4.68h3.05M11.5 25.4v9.37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.16 18.74a2.35 2.35 0 0 1 2.34-2.35h0m-2.34 0v6.21m-3.52 11.64a2.61 2.61 0 0 0 1.92.53h.53a1.56 1.56 0 0 0 1.55-1.55h0a1.56 1.56 0 0 0-1.55-1.56H32a1.54 1.54 0 0 1-1.55-1.55h0A1.55 1.55 0 0 1 32 28.56h.52a2.69 2.69 0 0 1 1.93.52m-3.03-7.66a2.35 2.35 0 0 1-2 1.18h0A2.34 2.34 0 0 1 27 20.26v-1.52a2.35 2.35 0 0 1 2.34-2.35h0a2.34 2.34 0 0 1 2.34 2.35v.76H27m-.36 7.13v7a1.17 1.17 0 0 0 1.17 1.17h.35m-2.75-6.24h2.46m-4.93 3.87a2.34 2.34 0 0 1-2.34 2.34h0a2.34 2.34 0 0 1-2.34-2.34V30.9a2.34 2.34 0 0 1 2.34-2.34h0a2.34 2.34 0 0 1 2.34 2.34m0 3.87v-6.21M11.5 13.23v6.27a3.11 3.11 0 0 0 6.21 0v-6.27m2.41 5.51a2.35 2.35 0 0 1 2.34-2.35h0a2.35 2.35 0 0 1 2.34 2.35v1.52a2.34 2.34 0 0 1-2.34 2.34h0a2.34 2.34 0 0 1-2.34-2.34m0 2.34v-9.37"/>`),
		g.Group(children),
	)
}