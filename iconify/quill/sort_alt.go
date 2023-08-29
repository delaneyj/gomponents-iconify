package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5.015v-1a1 1 0 0 0-1 1h1Zm8 1a1 1 0 1 0 0-2v2ZM4 13a1 1 0 1 0 2 0H4Zm23-7.985h1a1 1 0 0 0-1-1v1ZM26 13a1 1 0 1 0 2 0h-2Zm-7-8.985a1 1 0 1 0 0 2v-2Zm-13.293.278a1 1 0 0 0-1.414 1.414l1.414-1.414Zm22 1.414a1 1 0 0 0-1.414-1.414l1.414 1.414ZM11 26a1 1 0 1 0 0 2v-2Zm10 2a1 1 0 1 0 0-2v2ZM5 6.015h.015v-2H5v2Zm.015 0H13v-2H5.015v2ZM4 5.015V13h2V5.015H4Zm22 0V13h2V5.015h-2Zm1-1h-.015v2H27v-2Zm-.015 0H19v2h7.985v-2ZM16.707 15.293L5.722 4.308L4.308 5.722l10.985 10.985l1.414-1.414ZM5.722 4.308l-.015-.015l-1.414 1.414l.015.015l1.414-1.414Zm10.985 12.4L27.692 5.721l-1.414-1.414l-10.985 10.985l1.414 1.414ZM27.692 5.721l.015-.015l-1.414-1.414l-.015.015l1.414 1.414ZM17 27V16h-2v11h2Zm-6 1h5v-2h-5v2Zm5 0h5v-2h-5v2Z"/>`),
		g.Group(children),
	)
}