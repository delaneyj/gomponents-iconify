package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KniveForkCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M14.27 8.506c0 1.942 1.063 3.308 3 3.994V14a1 1 0 1 0 2 0V4.5a1 1 0 0 0-1.442-.897c-2.316 1.141-3.558 2.799-3.558 4.903Zm2 0c0-.82.319-1.535 1-2.161v3.971c-.695-.411-1-.999-1-1.81Z" clip-rule="evenodd"/><path d="M16.27 20.5v-5a2 2 0 1 1 4 0v5a2 2 0 0 1-4 0ZM6.283 5.45a1 1 0 1 1 1.998.1c-.08 1.603.002 2.682.2 3.158c.095.23.253.315.712.288a1 1 0 1 1 .114 1.997c-1.258.073-2.229-.446-2.674-1.519c-.343-.828-.444-2.142-.35-4.024Z"/><path d="M13.717 5.45a1 1 0 1 0-1.998.1c.08 1.603-.002 2.682-.2 3.158c-.096.23-.253.315-.712.288a1 1 0 1 0-.115 1.997c1.258.073 2.23-.446 2.675-1.519c.343-.828.444-2.142.35-4.024Z"/><path d="M9 5.5a1 1 0 0 1 2 0v9a1 1 0 1 1-2 0v-9Z"/><path d="M8 20.5v-5a2 2 0 1 1 4 0v5a2 2 0 1 1-4 0ZM4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}