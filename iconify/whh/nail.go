package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1021 1021q-16 16-113.5-53T776 865L249 338Q20 393 10 382q-10-9-10-22.5T10 336L336 10q10-10 23.5-10T382 10q11 10-44 239l527 527q34 34 103 131.5t53 113.5z"/>`),
		g.Group(children),
	)
}