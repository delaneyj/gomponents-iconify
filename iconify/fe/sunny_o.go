package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunnyO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feSunnyO0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSunnyO1" fill="currentColor" fill-rule="nonzero"><path id="feSunnyO2" d="M12 18a6 6 0 1 1 0-12a6 6 0 0 1 0 12Zm0-2a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM11 2h2v3h-2V2Zm-9 9h3v2H2v-2Zm17 0h3v2h-3v-2Zm-8 8h2v3h-2v-3Zm7.621-15l1.415 1.414l-2.122 2.122L16.5 6.12L18.621 4ZM16.5 17.414L17.914 16l2.122 2.121l-1.415 1.415l-2.121-2.122ZM6.121 16l1.415 1.414l-2.122 2.122L4 18.12L6.121 16ZM4 5.414L5.414 4l2.122 2.121L6.12 7.536L4 5.414Z"/></g></g>`),
		g.Group(children),
	)
}