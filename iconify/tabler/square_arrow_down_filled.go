package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareArrowDownFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M19 2a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3zm-7 5a1 1 0 0 0-1 1v5.585l-2.293-2.292l-.094-.083a1 1 0 0 0-1.32 1.497l4 4l.094.083l.092.064l.098.052l.11.044l.112.03l.126.017L12 17l.117-.007l.149-.029l.105-.035l.113-.054l.111-.071a.939.939 0 0 0 .112-.097l4-4l.083-.094a1 1 0 0 0-.083-1.32l-.094-.083a1 1 0 0 0-1.32.083L13 13.585V8l-.007-.117A1 1 0 0 0 12 7z"/></g>`),
		g.Group(children),
	)
}