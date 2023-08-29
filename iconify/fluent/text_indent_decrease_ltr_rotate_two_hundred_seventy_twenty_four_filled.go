package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentDecreaseLtrRotateTwoHundredSeventyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 15V6.5a1 1 0 0 1 1.993-.117L18 6.5V15a1 1 0 0 1-1.993.117L16 15Zm-4.707 6.707l-2-2a1 1 0 0 1 1.32-1.497l.094.083L12 19.586l1.293-1.293a1 1 0 0 1 1.497 1.32l-.083.094l-2 2a1 1 0 0 1-1.32.083l-.094-.083ZM11 15l-.001-11.5a1 1 0 0 1 1.993-.117L13 3.5V15a1 1 0 0 1-1.993.117L11 15Zm-5 0V6.5a1 1 0 0 1 1.993-.117L8 6.5V15a1 1 0 0 1-1.993.117L6 15Z"/>`),
		g.Group(children),
	)
}