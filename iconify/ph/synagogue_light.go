package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SynagogueLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M206 58.84V32a6 6 0 0 0-12 0v26.84A22 22 0 0 0 178 80v45.66l-44-25.14V72a6 6 0 0 0-12 0v28.52l-44 25.14V80a22 22 0 0 0-16-21.16V32a6 6 0 0 0-12 0v26.84A22 22 0 0 0 34 80v136a6 6 0 0 0 6 6h72a6 6 0 0 0 6-6v-40a10 10 0 0 1 20 0v40a6 6 0 0 0 6 6h72a6 6 0 0 0 6-6V80a22 22 0 0 0-16-21.16ZM200 70a10 10 0 0 1 10 10v26h-20V80a10 10 0 0 1 10-10ZM56 70a10 10 0 0 1 10 10v26H46V80a10 10 0 0 1 10-10Zm-10 48h20v92H46Zm82 36a22 22 0 0 0-22 22v34H78v-70.52l50-28.57l50 28.57V210h-28v-34a22 22 0 0 0-22-22Zm62 56v-92h20v92Z"/>`),
		g.Group(children),
	)
}