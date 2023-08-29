package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slashdot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M5 459L209 2h88L93 459H5zm374-69q0-31-21.5-52T306 317q-31 0-52 21t-21 52q0 30 21 51t52 21q30 0 51.5-21t21.5-51z"/>`),
		g.Group(children),
	)
}