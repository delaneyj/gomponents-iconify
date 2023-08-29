package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notification(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M649 8q19 0 36 7t30 20t19 30t7 36t-7 36t-19 29t-30 20t-36 8q-20 0-36-8t-30-20t-20-29t-7-36t7-36t20-30t30-20t36-7zm-93 255h93v463q0 10-7 17t-16 7H23q-10 0-16-7t-7-17V124q0-10 7-17t16-6h463v93H104q-11 0-11 11v440q0 12 11 12h440q12 0 12-12V263z"/>`),
		g.Group(children),
	)
}