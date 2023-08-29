package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dzone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M426 74q-28 0-35 27L231 48q0-8-6-14t-14-6t-14.5 6t-6.5 14t6 14L27 256q-2-1-6-1q-19 0-19 19q0 8 5.5 13.5T21 293q11 0 17-11l153 37q0 1-.5 3.5t-.5 3.5q0 14 10 24t24 10t24-10t10-24q0-7-1-10l117-48q5 5 13 5q7 0 12.5-5.5T405 255q0-10-7-15l19-94q1 0 4 .5t5 .5q15 0 25.5-11t10.5-26t-10.5-25.5T426 74zm-203-9l72 88l-73 30l-5-115zm-12 4l4 117l-33 13q0-1-6-7l28-124q2 1 7 1zm-27 140v-4l32-13l4 100q-2 0-5 1t-4 1l-33-70q6-6 6-15zm61 90q-9-7-19-7l-4-102l77-32l35 43zm59-143l88-35q0 2 5 11l-60 65zm86-41l-90 36l-73-89q2-3 4-9l159 53v9zM200 66l-28 124q-6-2-8-2q-8 0-14.5 6t-6.5 15q0 1 1 3v2L33 259q-1-1-2-1zM39 278q0-1 1-2v-2q0-6-3-9l110-45q5 9 17 9q4 0 10-2l33 70q-10 6-15 17zm216 34q-2-4-6-10l88-97l34 41q-2 6-2 9t2 9zm139-75q-1 0-3.5-.5t-3.5-.5q-9 0-13 6l-34-41l60-66q6 6 13 9z"/>`),
		g.Group(children),
	)
}