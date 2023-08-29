package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCommitDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M176 128a48 48 0 1 1-48-48a48 48 0 0 1 48 48Z" opacity=".2"/><path d="M248 120h-64.58a56 56 0 0 0-110.84 0H8a8 8 0 0 0 0 16h64.58a56 56 0 0 0 110.84 0H248a8 8 0 0 0 0-16Zm-120 48a40 40 0 1 1 40-40a40 40 0 0 1-40 40Z"/></g>`),
		g.Group(children),
	)
}