package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanTextTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.75 3A3.75 3.75 0 0 0 3 6.75V10a1 1 0 1 0 2 0V6.75C5 5.784 5.784 5 6.75 5H10a1 1 0 1 0 0-2H6.75ZM18 3a1 1 0 1 0 0 2h3.25c.966 0 1.75.784 1.75 1.75V10a1 1 0 1 0 2 0V6.75A3.75 3.75 0 0 0 21.25 3H18ZM5 18a1 1 0 1 0-2 0v3.25A3.75 3.75 0 0 0 6.75 25H10a1 1 0 1 0 0-2H6.75A1.75 1.75 0 0 1 5 21.25V18Zm20 0a1 1 0 1 0-2 0v3.25A1.75 1.75 0 0 1 21.25 23H18a1 1 0 1 0 0 2h3.25A3.75 3.75 0 0 0 25 21.25V18ZM9 8a1 1 0 0 0 0 2h10a1 1 0 1 0 0-2H9Zm0 5a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H9Zm-1 6a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1Z"/>`),
		g.Group(children),
	)
}