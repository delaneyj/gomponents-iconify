package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagEngland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="none"><path fill="#fff" d="M38 2.6H26C14.3 5 5 14.2 2.6 26v12C5 49.8 14.3 59 26 61.4h12C49.8 59 59 49.7 61.4 38V26C59 14.2 49.7 5 38 2.6z"/><path fill="#CE1124" d="M61.179 39A30.08 30.08 0 0 0 62 32c0-2.41-.284-4.754-.821-7H39V2.821A30.08 30.08 0 0 0 32 2c-2.41 0-4.754.284-7 .821V25H2.821A30.08 30.08 0 0 0 2 32c0 2.41.284 4.754.821 7H25v22.179A30.08 30.08 0 0 0 32 62c2.41 0 4.754-.284 7-.821V39h22.179z"/></g>`),
		g.Group(children),
	)
}