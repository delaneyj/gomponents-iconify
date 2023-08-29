package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M465 424q10 10 10 24t-10 24.5t-25 10.5t-24-10l-70-70q-10-10-10-24V194q0-15 10-25t24.5-10t24.5 10t10 25v170zM371 8q76 0 144 29t118 79t79 118t29 145t-29 144t-79 118t-118 80t-144.5 29T227 721t-118-80t-80-118T0 379t29-145t80-118t118-79T371 8zm0 649q57 0 108-22t89-60t59-88t21.5-108T627 270t-59-88t-89-59t-109-22t-107 22t-88 59t-60 88t-22 108.5T115 487t60 88t88 60t108 22z"/>`),
		g.Group(children),
	)
}