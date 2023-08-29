package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskListTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.707 3.293a1 1 0 0 0-1.414 0L4 4.586l-.293-.293a1 1 0 0 0-1.414 1.414l1 1a1 1 0 0 0 1.414 0l2-2a1 1 0 0 0 0-1.414Zm14.296 13.7H10L9.883 17A1 1 0 0 0 10 18.993h11.003l.117-.006a1 1 0 0 0-.117-1.994Zm0-5.993H10l-.117.007A1 1 0 0 0 10 13h11.003l.117-.007A1 1 0 0 0 21.003 11Zm0-6H10l-.117.007A1 1 0 0 0 10 7h11.003l.117-.007A1 1 0 0 0 21.003 5ZM6.707 16.293a1 1 0 0 0-1.414 0L4 17.586l-.293-.293a1 1 0 0 0-1.414 1.414l1 1a1 1 0 0 0 1.414 0l2-2a1 1 0 0 0 0-1.414Zm-1.414-6.5a1 1 0 0 1 1.414 1.414l-2 2a1 1 0 0 1-1.414 0l-1-1a1 1 0 1 1 1.414-1.414l.293.293l1.293-1.293Z"/>`),
		g.Group(children),
	)
}