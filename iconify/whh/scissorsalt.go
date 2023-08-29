package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scissorsalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m436.004 512l588 512q-56 0-181.5-61.5T587.004 816t-210-156q7 22 7 44q0 54-25 113t-71 101t-96 42q-80 0-136-56t-56-136t56-136t136-56q4 0 21 1q6-32 34-73q-15-31-21-57q-18 1-34 1q-80 0-136-56t-56-135.5t56-136t136-56.5q50 0 96 42t71 101t25 113q0 26-10 50q146-120 358.5-245t291.5-125zm-244-384q-53 0-90.5 37.5t-37.5 90.5t37.5 90.5t90.5 37.5h17q24 0 36-.5t29.5-3.5t25.5-9.5t14-19t6-31.5q0-36-16.5-80.5t-47.5-78t-64-33.5zm0 512q-53 0-90.5 37.5t-37.5 90.5t37.5 90.5t90.5 37.5q33 0 64-33.5t47.5-78t16.5-80.5q0-19-6-31.5t-14-19t-25.5-9.5t-29.5-3.5t-36-.5h-17z"/>`),
		g.Group(children),
	)
}