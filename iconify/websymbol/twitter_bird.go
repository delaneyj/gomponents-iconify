package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitterBird(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1003 501q-41 46-114 46h-6q-37 152-177.5 244T401 883q-121 0-225.5-53T0 679q75 72 181 72q113 0 191-83q-6 1-12 1l-36-8q-24-12-24-35q0-27 41-45q-10 1-19 1q-29 0-54-12q-33-16-51-51q22-25 69-28q-98-24-112-102q27-8 53-8h8q-37-20-62-51.5T149 260l1-8q155 59 257 116q30 17 76 63q32-87 67.5-150.5T640 182q-1 16-15 31q33-32 78-38q-3 23-53 41q7-2 25.5-9.5t33-11.5t25.5-4q15 0 15 11q0 8-16.5 16T690 233.5t-28 8.5q10-1 19-1q69 0 123 52q60 60 76 143q18 6 36 6q50 0 83-19q-13 30-43.5 47T889 489q34 15 83 15q16 0 31-3z"/>`),
		g.Group(children),
	)
}