package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M714 190q11 0 18 7t5 18l-42 428q-2 10-12 10H55q-11 0-11-10L1 215q-1-10 6-18t17-7h690zm-20-46H45V28q0-10 7-17t16-6h185q10 0 17 6t7 17v23h393q10 0 17 6t7 17v70z"/>`),
		g.Group(children),
	)
}