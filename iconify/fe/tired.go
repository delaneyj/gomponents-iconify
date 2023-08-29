package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tired(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTired0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTired1" fill="currentColor" fill-rule="nonzero"><path id="feTired2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10ZM8.2 10L7 8.8l.8-.8L9 9.2L10.2 8l.8.8L9.8 10l1.2 1.2l-.8.8L9 10.8L7.8 12l-.8-.8L8.2 10Zm6 0L13 8.8l.8-.8L15 9.2L16.2 8l.8.8l-1.2 1.2l1.2 1.2l-.8.8l-1.2-1.2l-1.2 1.2l-.8-.8l1.2-1.2ZM12 17a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`),
		g.Group(children),
	)
}