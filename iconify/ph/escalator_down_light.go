package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorDownLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M171.76 84.24a6 6 0 0 1 8.48-8.48L194 89.51V48a6 6 0 0 1 12 0v41.51l13.76-13.75a6 6 0 0 1 8.48 8.48l-24 24a6 6 0 0 1-8.48 0ZM238 160v40a14 14 0 0 1-14 14h-52.5a14 14 0 0 1-10.29-4.51L70 110.64a2 2 0 0 0-1.5-.64H32a14 14 0 0 1-14-14V56a14 14 0 0 1 14-14h52.5a14 14 0 0 1 10.29 4.51L186 145.36a2 2 0 0 0 1.47.64H224a14 14 0 0 1 14 14Zm-12 0a2 2 0 0 0-2-2h-36.5a14 14 0 0 1-10.29-4.51L86 54.64a2 2 0 0 0-1.5-.64H32a2 2 0 0 0-2 2v40a2 2 0 0 0 2 2h36.5a14 14 0 0 1 10.29 4.51L170 201.36a2 2 0 0 0 1.47.64H224a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}