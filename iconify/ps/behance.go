package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Behance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M418 89H302V61h116v28zM225 207q12 17 12 42q0 23-13 46q-12 16-20 22q-16 11-33 14q-16 4-40 4H2V49h138q53 0 74 30q13 17 13 44q0 28-13 42q-9 11-22 16q23 9 33 26zM68 162h60q17 0 31-7q11-8 11-26q0-20-15-26q-15-5-34-5H68v64zm108 83q0-22-18-31q-12-5-29-5H68v77h60q17 0 29-5q19-10 19-36zm285-47q2 20 1 41H313q1 31 21 43q13 8 30 8q19 0 30-10q6-5 11-14h55q-3 18-20 38q-29 29-78 29q-41 0-72-25t-31-82q0-54 28-82t73-28q26 0 49 9q22 11 35 31q12 16 17 42zm-54 5q-2-22-15-32q-13-11-32-11q-20 0-32 12q-11 11-14 31h93z"/>`),
		g.Group(children),
	)
}