package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Videocamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M726 146q5-2 10 1t5 10v244q0 6-5 9t-10 2l-100-40V186zM510 1q19 0 33 13t13 33v464q0 19-13 32t-33 14H46q-19 0-32-14T0 511V47q0-19 14-33T46 1h464zM278 441q34 0 64-13t51-35t35-52t12-62q0-34-12-64t-35-51t-51-35t-64-12t-63 12t-52 35t-34 51t-13 64q0 33 13 62t34 52t52 35t63 13zm0-255q19 0 36 7t30 20t20 30t7 36t-7 36t-20 29t-30 20t-36 8t-36-8t-29-20t-20-29t-8-36t8-36t20-30t29-20t36-7z"/>`),
		g.Group(children),
	)
}