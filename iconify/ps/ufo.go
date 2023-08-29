package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ufo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m516 211l6-10l-2-7q0-4-8.5-18.5T471 136t-81-40q-8-41-44-68.5T264 0t-82 27.5T138 96Q35 128 9 193l-1 1l-2 9l4 8l2 2H8q0 58 41 103q66 68 215 68t215-70q41-43 41-103h-4zM264 43q33 0 57.5 19.5T349 111l-4 17h2q-9 43-83 43t-83-43h2l-4-17q3-29 27.5-48.5T264 43zm-126 98q2 12 13 28l-30 29q-14 16 0 30q6 7 15 7t15-7l32-32q27 14 58 17v43q0 21 21 21t21-21v-45q34-3 58-17l32 32q7 7 15 7q6 0 15-7q14-14 0-30l-30-30q8-11 13-27q61 22 85 57q-17 28-68 54.5T264 277q-61 0-109.5-16.5t-68-31.5T55 198q24-36 83-57zm126 200q-131 0-183-55q-7-7-15-19q85 53 198 53t198-53q-8 12-15 19q-52 55-183 55z"/>`),
		g.Group(children),
	)
}