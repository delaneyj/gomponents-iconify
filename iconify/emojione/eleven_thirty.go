package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevenThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#62727a"/><g fill="#fff"><path d="M30 26h4v32h-4z"/><path d="m35.4 37l-3.8 1l-6.2-23l3.8-1z"/><circle cx="32" cy="32" r="4"/></g><circle cx="32" cy="32" r="3" fill="#62727a"/>`),
		g.Group(children),
	)
}