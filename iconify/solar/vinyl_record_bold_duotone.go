package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylRecordBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4.929 19.071c3.905 3.905 10.237 3.905 14.142 0c3.905-3.905 3.905-10.237 0-14.142c-3.905-3.905-10.237-3.905-14.142 0c-3.905 3.905-3.905 10.237 0 14.142Z" opacity=".5"/><path d="M17.127 6.873a.75.75 0 1 0-1.061 1.06a5.75 5.75 0 0 1 0 8.132a.75.75 0 1 0 1.06 1.061a7.25 7.25 0 0 0 0-10.253ZM7.934 7.934a.75.75 0 0 0-1.06-1.061a7.25 7.25 0 0 0 0 10.253a.75.75 0 0 0 1.06-1.06a5.75 5.75 0 0 1 0-8.132Z"/><path fill-rule="evenodd" d="M9.348 9.348a3.75 3.75 0 1 1 5.304 5.303a3.75 3.75 0 0 1-5.304-5.303Zm1.061 1.06a2.25 2.25 0 1 1 3.182 3.183a2.25 2.25 0 0 1-3.182-3.182Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}