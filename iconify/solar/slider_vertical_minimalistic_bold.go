package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SliderVerticalMinimalisticBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5.846 8c0-2.828 0-4.243.901-5.121C7.65 2 9.1 2 12 2c2.901 0 4.351 0 5.253.879c.9.878.9 2.293.9 5.121v8c0 2.828 0 4.243-.9 5.121C16.35 22 14.9 22 12 22c-2.901 0-4.351 0-5.253-.879c-.9-.878-.9-2.293-.9-5.121V8Z"/><path fill-rule="evenodd" d="M2.77 3.75a.76.76 0 0 1 .768.75v15a.76.76 0 0 1-.769.75A.76.76 0 0 1 2 19.5v-15a.76.76 0 0 1 .77-.75Zm18.46 0a.76.76 0 0 1 .77.75v15a.76.76 0 0 1-.77.75a.76.76 0 0 1-.768-.75v-15a.76.76 0 0 1 .769-.75Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}