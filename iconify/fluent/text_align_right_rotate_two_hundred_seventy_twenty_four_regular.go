package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignRightRotateTwoHundredSeventyTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.75 19a.75.75 0 0 1-.75-.75V2.75a.75.75 0 0 1 1.5 0v15.5a.75.75 0 0 1-.75.75Zm13-5a.75.75 0 0 1-.75-.75V2.75a.75.75 0 0 1 1.5 0v10.5a.75.75 0 0 1-.75.75Zm-7.25 7.25a.75.75 0 0 0 1.5 0V2.75a.75.75 0 0 0-1.5 0v18.5Z"/>`),
		g.Group(children),
	)
}