package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Behance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M333 250q38 0 58 20t28 46q10 29 9 66q0 31-7 54t-18 39t-26 28t-30 19q-36 17-81 18H0V7h266q30 0 57 8t46 27t30 44t11 63t-12 56t-27 30q-17 11-38 15zm-216-30h140q18 0 32-11t13-52q0-20-5-32t-13-17t-18-6H117v118zm141 225q8 0 19-3t20-9t17-21t6-36q0-28-8-43t-19-21q-12-8-27-9H117v142h141zm387-303q39 0 67 10t48 26t32 37t19 43q16 51 11 116H558q0 38 14 56t31 27q19 10 45 9q32 0 49-9t24-19q9-12 9-29h90q0 29-7 50t-20 36t-29 26t-32 16q-39 15-87 14q-39 0-68-9t-51-23t-35-33t-23-37q-20-45-20-102q3-57 25-102q9-19 24-38t35-33t48-23t65-9zm80 161q0-23-9-41q-8-16-24-29t-45.5-13t-46.5 13t-28 30t-15 40h168zm21-199H535V38h211v66z"/>`),
		g.Group(children),
	)
}