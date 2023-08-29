package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Movie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M136 488h359.985V24H16v464h120ZM408 56h55.985v40H408Zm0 72h55.985v40H408Zm0 72h55.985v40H408Zm0 72h55.985v40H408Zm0 72h55.985v40H408Zm0 72h55.985v40H408ZM136 200V56h239.985v184H136Zm0 216V272h239.985v184H136ZM48 56h56v40H48Zm0 72h56v40H48Zm0 72h56v40H48Zm0 72h56v40H48Zm0 72h56v40H48Zm0 72h56v40H48Z"/>`),
		g.Group(children),
	)
}