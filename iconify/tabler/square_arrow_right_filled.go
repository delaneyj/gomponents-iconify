package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareArrowRightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M19 2a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3zm-6.387 5.21a1 1 0 0 0-1.32.083l-.083.094a1 1 0 0 0 .083 1.32L13.585 11H8l-.117.007A1 1 0 0 0 8 13h5.585l-2.292 2.293l-.083.094a1 1 0 0 0 1.497 1.32l4-4l.073-.082l.074-.104l.052-.098l.044-.11l.03-.112l.017-.126L17 12l-.007-.118l-.029-.148l-.035-.105l-.054-.113l-.071-.111a1.008 1.008 0 0 0-.097-.112l-4-4z"/></g>`),
		g.Group(children),
	)
}