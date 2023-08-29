package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingImageRightParagraphImageTextAlignmentWrapRightFormatting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 1H5M.5 4H5M.5 7H5M.5 13h13m-13-3h13"/><rect width="6" height="6" x="7.5" y="1" rx=".5" transform="rotate(-180 10.5 4)"/></g>`),
		g.Group(children),
	)
}