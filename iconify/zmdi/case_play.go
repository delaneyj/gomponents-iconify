package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CasePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M299 88h128v277q0 18-12.5 30.5T384 408H43q-18 0-30.5-12.5T0 365V88h128V45q0-17 12.5-29.5T171 3h85q18 0 30.5 12.5T299 45v43zM171 45v43h85V45h-85zm-22 299l160-107l-160-85v192z"/>`),
		g.Group(children),
	)
}