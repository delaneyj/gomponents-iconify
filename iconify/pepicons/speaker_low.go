package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m10.334 1.754l-4.68 4.178H2a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3.535l4.796 4.312c.644.578 1.669.122 1.669-.744v-15c0-.864-1.021-1.321-1.666-.746ZM6.666 7.709L10 4.733v10.523l-3.126-2.81A1 1 0 0 0 6 11.932H3v-4h2.751c.302.079.642.02.915-.223Z" clip-rule="evenodd"/><path d="M15.489 13.069a.75.75 0 0 1-.978-1.138c.034-.029.067-.06.1-.094c.138-.143.261-.324.362-.536A3.05 3.05 0 0 0 15.25 10c0-.754-.25-1.433-.639-1.837a1.464 1.464 0 0 0-.1-.094a.75.75 0 1 1 .978-1.138c.07.06.137.124.202.191c.67.696 1.059 1.75 1.059 2.878c0 .696-.147 1.366-.423 1.945c-.168.355-.383.67-.636.933a2.964 2.964 0 0 1-.202.19Z"/></g>`),
		g.Group(children),
	)
}