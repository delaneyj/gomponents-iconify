package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylRecordBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.929 19.071c3.905 3.905 10.237 3.905 14.142 0c3.905-3.905 3.905-10.237 0-14.142c-3.905-3.905-10.237-3.905-14.142 0c-3.905 3.905-3.905 10.237 0 14.142ZM17.126 6.873a.75.75 0 1 0-1.06 1.061a5.75 5.75 0 0 1 0 8.132a.75.75 0 1 0 1.06 1.06a7.25 7.25 0 0 0 0-10.253ZM9.349 9.348a3.75 3.75 0 1 1 5.304 5.304a3.75 3.75 0 0 1-5.304-5.304Zm1.061 1.061a2.25 2.25 0 1 1 3.182 3.182a2.25 2.25 0 0 1-3.182-3.182ZM7.934 7.934a.75.75 0 0 0-1.06-1.06a7.25 7.25 0 0 0 0 10.252a.75.75 0 0 0 1.06-1.06a5.75 5.75 0 0 1 0-8.132Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}