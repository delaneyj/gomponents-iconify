package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalSee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.5q-1.05 0-1.775-.725T9.5 14q0-1.05.725-1.775T12 11.5h.05q.55.775 1.163 1.5t1.262 1.4q-.125.9-.825 1.5t-1.65.6ZM4 22q-.825 0-1.413-.588T2 20V8q0-.825.588-1.413T4 6h3.15L9 4h1.25q-.125.475-.188.975T10 6q0 .95.25 1.85t.675 1.775q-1.5.35-2.463 1.563T7.5 14q0 1.875 1.313 3.188T12 18.5q1.35 0 2.438-.713t1.637-1.862q.5.475.988.9t.937.8q1.05-.875 2.05-1.788T22 13.95V20q0 .825-.588 1.413T20 22H4Zm14-7q3.025-2.575 4.513-4.775T24 6.15q0-2.825-1.813-4.487T18 0q-2.375 0-4.188 1.663T12 6.15q0 1.875 1.488 4.075T18 15Zm-1.85-6l.7-2.275L15 5.25h2.275L18 3l.725 2.25H21l-1.85 1.475l.7 2.275L18 7.6L16.15 9Z"/>`),
		g.Group(children),
	)
}