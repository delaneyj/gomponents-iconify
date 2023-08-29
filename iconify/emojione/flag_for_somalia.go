package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSomalia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#42ade2"/><path fill="#fff" d="m32 39.2l9.9 7.1l-3.8-11.5l9.9-7.1H35.8L32 16.3l-3.8 11.4H16l9.8 7.1l-3.7 11.5z"/>`),
		g.Group(children),
	)
}