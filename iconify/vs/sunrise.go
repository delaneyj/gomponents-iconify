package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1565 1216h-93q0-157-77-289.5T1185.5 717T896 640t-289.5 77T397 926.5T320 1216h-93L47 968q-9-13-4-29q4-15 20-20l292-97V516q2-18 14-26q14-9 29-4l292 94l180-248q9-12 26-12t26 12l180 248l292-94q15-5 29 4q13 10 13 26v306l293 97q6 2 13 8t7 12q4 15-4 29l-181 248h1zM55 1312q-21 0-38 16.5T0 1367v50q0 22 17 38.5t38 16.5h1682q21 0 38-16.5t17-38.5v-50q0-22-17-38.5t-38-16.5H55zm841-576q-130 0-240.5 64.5t-175 175T416 1216h960q0-130-64.5-240.5t-175-175T896 736z"/>`),
		g.Group(children),
	)
}