package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSubbed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M29 9v7m0-7c0-1-1-2-2-2H5C4 7 3 8 3 9m26 0l-11.862 8.212a2 2 0 0 1-2.276 0L3 9m13 16H5c-1 0-2-1-2-2V9m16.5 14l3 3l6-6"/>`),
		g.Group(children),
	)
}