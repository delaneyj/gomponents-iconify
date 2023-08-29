package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v7h2V5h16v5h2V3H6zm2 9l2 4l-2 4h2l1-2l1 2h2l-2-4l2-4h-2l-1 2l-1-2H8zm7 0v8h4v-2h-2v-6h-2zm7.5 0c-1.383 0-2.5 1.117-2.5 2.5s1.117 2.5 2.5 2.5c.217 0 .5.283.5.5c0 .3-.12.5-.5.5c-.368 0-.424-.08-.438-.094c-.013-.013-.062-.08-.062-.312h-2c0 .566.163 1.2.625 1.687c.462.488 1.143.72 1.875.72c1.42 0 2.5-1.2 2.5-2.5c0-1.383-1.117-2.5-2.5-2.5c-.217 0-.5-.283-.5-.5c0-.217.283-.5.5-.5c.267 0 .348.063.406.125c.06.062.094.17.094.28h2c0-.587-.215-1.192-.656-1.655c-.442-.463-1.11-.75-1.844-.75zM6 22v7h20v-7h-2v5H8v-5H6z"/>`),
		g.Group(children),
	)
}