package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterVintage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.05 22q-1.425 0-2.525-.9T8.15 18.8q-1.325.425-2.675-.05T3.4 17.1q-.75-1.2-.55-2.663T4.15 12q-1.05-.95-1.263-2.35t.463-2.6q.675-1.2 2.038-1.737T8.1 5.2q.275-1.4 1.375-2.3T12 2q1.425 0 2.525.9T15.9 5.2q1.4-.425 2.713.075T20.65 7.05q.675 1.25.488 2.613T19.85 12q1.1.975 1.313 2.413T20.7 17.1q-.725 1.275-2.037 1.7t-2.713 0q-.275 1.4-1.375 2.3t-2.525.9ZM12 16q1.65 0 2.825-1.175T16 12q0-1.65-1.175-2.825T12 8q-1.65 0-2.825 1.175T8 12q0 1.65 1.175 2.825T12 16Z"/>`),
		g.Group(children),
	)
}