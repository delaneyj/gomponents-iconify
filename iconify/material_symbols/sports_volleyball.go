package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsVolleyball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.6 9.2L13 4.225V2.05q3.075.3 5.4 2.263T21.6 9.2ZM6.75 13.9V3.5q.975-.575 2.05-.95t2.225-.5v9.375L6.75 13.9ZM2.9 16.125q-.45-.95-.675-1.988T2 12q0-1.925.7-3.688T4.75 5.15v9.875l-1.85 1.1ZM7.3 20.8q-1.05-.5-1.9-1.25t-1.5-1.7l8.125-4.7l4.275 2.475l-9 5.175Zm4.725 1.2q-.575 0-1.175-.063t-1.15-.212l8.6-4.95l1.85 1.025q-1.425 1.95-3.562 3.075T12.025 22Zm9.1-5.9L13 11.425V6.5l9 5.225q0 1.15-.2 2.25t-.675 2.125Z"/>`),
		g.Group(children),
	)
}