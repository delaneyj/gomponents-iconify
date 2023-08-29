package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M162 483q18 15 42 15q32 0 51-24l163-194q26-32 29-75l11-143q0-24-15-36l-17-13q-19-14-38-8L247 41q-42 9-68 43L17 278q-19 23-15 49q3 28 23 45zM49 303l162-194q15-19 47-30l140-36l17 15l-10 143q-3 29-22 51L221 446q-6 8-15.5 8.5T189 449L53 333q-9-9-9-15q0-7 5-15zm311-183q0 9-6.5 15t-14.5 6q-9 0-15.5-6t-6.5-15t6.5-15t15.5-6q8 0 14.5 6t6.5 15z"/>`),
		g.Group(children),
	)
}