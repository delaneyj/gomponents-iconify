package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMedal0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMedal1" fill="currentColor"><path id="feMedal2" d="M12 22a7 7 0 1 1 0-14a7 7 0 0 1 0 14Zm0-2a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm0-12a1 1 0 1 1 0-2a1 1 0 0 1 0 2ZM9 2h2v4c-1.1 0-2-.9-2-2V2Zm4 0h2v2c0 1.1-.9 2-2 2V2Z"/></g></g>`),
		g.Group(children),
	)
}