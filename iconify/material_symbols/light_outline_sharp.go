package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-1.65 0-2.825-1.175T8 17H3.225q-.1-.4-.163-.913T3 15q.05-3.5 2.325-6.038T11 6.05V3h2v3.05q3.4.375 5.675 2.913T21 15q0 .575-.063 1.088t-.162.912H16q0 1.65-1.175 2.825T12 21Zm-7-6h14q0-2.9-2.05-4.95T12 8q-2.9 0-4.95 2.05T5 15Zm7 4q.825 0 1.413-.588T14 17h-4q0 .825.588 1.413T12 19Zm0-2Z"/>`),
		g.Group(children),
	)
}