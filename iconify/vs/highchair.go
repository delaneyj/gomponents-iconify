package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Highchair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M768 896V512h229q11 0 19-8t8-19v-74q0-11-8-19t-19-8H256V24q0-10-7-17t-17-7h-80q-10 0-17 7t-7 17v788q0 34 17 57.5t47 26.5L0 1792h130l94-384h464l77 384h131L704 896h64zm-104 384H248l72-384h256z"/>`),
		g.Group(children),
	)
}