package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Genetics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 23v-1q0-3.475 1.45-5.638t4-4.362q-2.55-2.2-4-4.362T5 2V1h2v1q0 .275.013.513T7.05 3h9.9q.025-.25.038-.488T17 2V1h2v1q0 3.475-1.45 5.638t-4 4.362q2.55 2.2 4 4.362T19 22v1h-2v-1q0-.275-.013-.513T16.95 21h-9.9q-.025.25-.037.488T7 22v1H5ZM8.45 7h7.1q.325-.475.563-.95T16.55 5h-9.1q.2.55.437 1.038T8.45 7ZM12 10.7q.5-.425.975-.85t.9-.85h-3.75q.425.425.9.85t.975.85ZM10.125 15h3.75q-.425-.425-.9-.85T12 13.3q-.5.425-.975.85t-.9.85ZM7.45 19h9.1q-.2-.55-.438-1.038T15.55 17h-7.1q-.325.475-.562.95T7.45 19Z"/>`),
		g.Group(children),
	)
}