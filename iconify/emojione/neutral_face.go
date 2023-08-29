package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeutralFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><g fill="#664e27"><circle cx="20.5" cy="27.6" r="5"/><circle cx="43.5" cy="27.6" r="5"/><path d="M38.9 48H25.1c-1.5 0-1.5-4 0-4h13.7c1.6 0 1.6 4 .1 4"/></g>`),
		g.Group(children),
	)
}