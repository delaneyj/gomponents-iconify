package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marriott(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.802 11.083l-1.178 2.41c-.8 1.425-1.931 3.167-3.646 3.603c-.668.232-1.255.023-1.9-.023L0 20.476a1.626 1.626 0 0 0 .59.386c3.647 1.39 5.122-.1 8.722-8.238l3.403 7.249h4.53l-2.14-4.893l1.213-2.53l3.345 7.311l4.337.027l-7.59-16.677l-3.475 1.738l2.738 6.222l-1.201 2.445L9.45 2.678l-3.7 1.877Z"/>`),
		g.Group(children),
	)
}