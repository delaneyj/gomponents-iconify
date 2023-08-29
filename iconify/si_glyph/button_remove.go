package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.021.114C3.627.114.052 3.668.052 8.036s3.575 7.922 7.969 7.922c4.395 0 7.969-3.554 7.969-7.922S12.415.114 8.021.114zM8.013 14.07c-3.329 0-6.04-2.724-6.04-6.071c0-3.347 2.711-6.071 6.04-6.071s6.039 2.725 6.039 6.071c0 3.348-2.71 6.071-6.039 6.071z"/><path d="M4.813 8.982c-.265 0-.782-.298-.782-.986c0-.654.438-.959.732-.959l6.453-.002c.265 0 .783.299.783.987c0 .653-.438.958-.732.958l-6.454.002z"/></g>`),
		g.Group(children),
	)
}