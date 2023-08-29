package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.987 6.042c-1.11 0-2.011.88-2.011 1.968c0 1.084.9 1.964 2.011 1.964c1.108 0 2.009-.88 2.009-1.964c0-1.088-.9-1.968-2.009-1.968z"/><path d="m10.979 12.963l1.95-.008v-2.018h2.507a7.793 7.793 0 0 0 .565-2.896c0-4.396-3.582-7.958-8-7.958s-8 3.562-8 7.958c0 4.396 3.582 7.959 8 7.959a8.014 8.014 0 0 0 3.333-.73l.667-.301l-1.022-2.006zm1.444-10.279l.846.846l-1.787 1.787l-.845-.845l1.786-1.788zM3.548 13.25l-.845-.846l1.788-1.787l.845.845l-1.788 1.788zM8 11a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/><path d="m14.521 15.958l1.373-1.937l-.921.003v-1.962h-.944v1.964l-.982.004l1.474 1.928z"/></g>`),
		g.Group(children),
	)
}