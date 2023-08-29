package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emailforward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1017 989L667 640l350-349q7 14 7 29v640q0 15-7 29zM63 256h135q23-110 111-183T512 0t203 73t111 183h135L512 705zm495 248l146-184l-146-185q-10-8-23-8t-23 8v121H352q-13 0-22.5 9.5T320 288v64q0 13 9.5 22.5T352 384h160v120q10 8 23 8t23-8zM7 989q-7-14-7-29V320q0-15 7-29l350 349zm466-233q6 6 15.5 9t16.5 3h7q26 1 39-12l71-71l339 339H63l339-339z"/>`),
		g.Group(children),
	)
}