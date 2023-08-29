package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCircleO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCheckCircleO0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCheckCircleO1" fill="currentColor" fill-rule="nonzero"><path id="feCheckCircleO2" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm0 2C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10ZM8 10l-2 2l5 5l7-7l-2-2l-5 5l-3-3Z"/></g></g>`),
		g.Group(children),
	)
}