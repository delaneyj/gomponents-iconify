package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func World(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M371 8q76 0 144 29t118 79t79 118t29 145t-29 144t-79 118t-118 80t-144.5 29T227 721t-118-80t-80-118T0 379t29-145t80-118t118-79T371 8zM93 379q0 49 16 93t45 80t68 61t85 36q-13-19-9-35q5-20 25-40t44-29.5t41-23.5t17-23t-23-24t-58-22q-17-5-32-14t-28-20t-21-23t-12-24q-5-24-18-31t-27.5 7.5T178 372t-19 10t2-20t21-49t21-53t2-28t-21-9t-37-3q-3 0-4-1q-24 35-37 75t-13 85zm512 150q20-32 32-70t12-80q0-36-9-70t-25-63q-1 3-5 9q-9 14-26 34t-31.5 34.5T531 338t-2-14t-8-21t-38-2t-47 15t-34 20t-10 23t10 23t27 10t31-10t28-10t18 10t2 33t-13 48q-10 24-3 38t25 9t43 0t45 19z"/>`),
		g.Group(children),
	)
}