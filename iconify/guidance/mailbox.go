package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mailbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M11 21.5v-7.174C11 11.27 9.058 8.5 6 8.5s-5 2.769-5 5.826V21.5h10Zm0 0h.5l11.5-2v-7.541c0-2.845-2.35-4.816-5.163-4.934a3.554 3.554 0 0 0-.534.023L16 7.192m-2 .227c-2.382.275-6.484.8-8.6 1.118M14 16.5V4.75m0 0v-3L19 1s.5 1.5 0 3l-5 .75ZM3 14s1.5.5 3 .5s3-.5 3-.5"/>`),
		g.Group(children),
	)
}