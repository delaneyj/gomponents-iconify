package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMail0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMail1" fill="currentColor" fill-rule="nonzero"><path id="feMail2" d="M4 10v8h16v-8l-8 3l-8-3Zm0-4v2l8 3l8-3V6H4Zm0-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Z"/></g></g>`),
		g.Group(children),
	)
}