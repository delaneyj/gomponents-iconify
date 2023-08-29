package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 512H64q-26 0-45-19T0 447.5t19-45T64 384h192q27 0 45.5-19t18.5-45.5t-18.5-45t-45-18.5t-45.5-19t-19-45.5t19-45t45-18.5q6 0 13 1q75 5 127 60t52 131q0 31-12 64h332q53 0 90.5-37.5T896 256t-37.5-90.5T768 128t-90.5 37.5T640 256q0 26-18.5 45t-45 19t-45.5-19t-19-45q0-106 75-181T768 0t181 75t75 181t-75 181t-181 75zM64 576h320q4 0 11 1q11-1 21-1q93 0 158.5 65.5T640 800q0 48-13.5 87.5T584 958t-79.5 48.5T384 1024q-27 0-45.5-19T320 959.5t19-45t45-18.5q31 0 53.5-3t36.5-10.5t22-14t11.5-20.5t4-22.5t.5-25.5q0-40-28-68t-68-28H64q-26 0-45-19T0 639.5t18.5-45T64 576z"/>`),
		g.Group(children),
	)
}