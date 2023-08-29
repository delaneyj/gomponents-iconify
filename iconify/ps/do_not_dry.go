package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotDry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M471 32q2-2 4-2q-11-30-40-30H93Q64 0 53 28q1 0 2 1t2 1l36 32V43h342v19zm-36 309l42 34V137l-42 34v170zM57 480q-1 0-2 1t-2 1q11 30 40 30h342q29 0 40-30q-1 0-2-1t-2-1l-36-30v19H93v-19zm36-309l-42-34v238l42-34V171zM17 431q-7 5-8 14t5 16q4 8 15 8q9 0 13-4l9-6l42-34l171-139l171 139l42 34l9 6q4 4 13 4q10 0 17-8q5-7 4.5-16.5T514 431l-37-28l-42-34l-137-113l137-113l42-34l34-28q7-5 8-14t-5-16q-4-8-15-8q-9 0-13 4l-9 6l-42 34l-171 139L93 90L51 53l-9-6q-4-4-13-4q-10 0-17 8q-11 17 2 30l37 28l42 34l137 113L93 369l-42 34z"/>`),
		g.Group(children),
	)
}