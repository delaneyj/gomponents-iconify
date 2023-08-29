package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareArrowUpFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M19 2a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3zm-7 5l-.09.004l-.058.007l-.118.025l-.105.035l-.113.054l-.111.071a1.008 1.008 0 0 0-.112.097l-4 4l-.083.094a1 1 0 0 0 .083 1.32l.094.083a1 1 0 0 0 1.32-.083L11 10.415V16l.007.117A1 1 0 0 0 13 16v-5.585l2.293 2.292l.094.083a1 1 0 0 0 1.32-1.497l-4-4l-.082-.073l-.104-.074l-.098-.052l-.11-.044l-.112-.03l-.126-.017L12 7z"/></g>`),
		g.Group(children),
	)
}