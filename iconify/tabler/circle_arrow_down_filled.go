package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleArrowDownFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M17 3.34a10 10 0 1 1-14.995 8.984L2 12l.005-.324A10 10 0 0 1 17 3.34zM12 7a1 1 0 0 0-1 1v5.585l-2.293-2.292l-.094-.083a1 1 0 0 0-1.32 1.497l4 4c.028.028.057.054.094.083l.092.064l.098.052l.081.034l.113.034l.112.02L12 17l.115-.007l.114-.02l.142-.044l.113-.054l.111-.071a.939.939 0 0 0 .112-.097l4-4l.083-.094a1 1 0 0 0-1.497-1.32L13 13.584V8l-.007-.117A1 1 0 0 0 12 7z"/></g>`),
		g.Group(children),
	)
}