package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm560 1097l265 321q19 24 29 35t30.5 25.5t38.5 14.5h159q27 0 51-15t24-40q0-20-34-60l-399-482l399-482q34-40 34-60q0-25-24-40t-51-15h-159q-18 0-38.5 14.5T1190 339t-29 35L896 695L631 374q-19-24-29-35t-30.5-25.5T533 299H374q-27 0-51 15t-24 40q0 20 34 60l399 482l-399 482q-34 40-34 60q0 25 24 40t51 15h159q18 0 38.5-14.5T602 1453t29-35z"/>`),
		g.Group(children),
	)
}