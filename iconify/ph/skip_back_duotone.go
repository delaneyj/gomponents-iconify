package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBackDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 47.88v160.24a8 8 0 0 1-12.19 6.65L59.7 134.65a7.83 7.83 0 0 1 0-13.3l128.11-80.12A8 8 0 0 1 200 47.88Z" opacity=".2"/><path d="M199.81 34a16 16 0 0 0-16.24.43L64 109.23V40a8 8 0 0 0-16 0v176a8 8 0 0 0 16 0v-69.23l119.57 74.78A15.95 15.95 0 0 0 208 208.12V47.88A15.86 15.86 0 0 0 199.81 34ZM192 208L64.16 128L192 48.07Z"/></g>`),
		g.Group(children),
	)
}