package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatAddTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 10a8 8 0 1 0-16 0l.007.346l.026.382a7.95 7.95 0 0 0 .829 2.887l.063.12l-.91 3.644l-.014.083v.082a.5.5 0 0 0 .62.441l3.645-.91l.12.064a7.968 7.968 0 0 0 3.868.857a5.5 5.5 0 0 1 7.742-7.742c.003-.084.004-.17.004-.254Zm1 4.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-4-2a.5.5 0 0 0-1 0V14h-1.5a.5.5 0 0 0 0 1H14v1.5a.5.5 0 0 0 1 0V15h1.5a.5.5 0 0 0 0-1H15v-1.5Z"/>`),
		g.Group(children),
	)
}