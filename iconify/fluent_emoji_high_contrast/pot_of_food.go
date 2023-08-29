package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PotOfFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M7 19H6a2 2 0 1 1 0-4h1v4Zm18 0h1a2 2 0 1 0 0-4h-1v4Zm-7-1a2 2 0 1 1 4 0h-4Zm-7 1a1 1 0 1 1 2 0h-2Zm7-7a2 2 0 1 1 4 0h-4Zm0 0h-2a1 1 0 1 1 2 0Zm-2 7a2 2 0 0 0-2 2h4a2 2 0 0 0-2-2Zm-.25-10a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm-5 7a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0ZM21 14.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm-6.689-2.024a1.75 1.75 0 1 0-3.5 0a4.495 4.495 0 0 0 4.495 4.495a1.75 1.75 0 1 0 0-3.5a.995.995 0 0 1-.995-.995Z"/><path d="M26.829 22.943C25.91 28.091 21.412 32 16 32c-5.412 0-9.911-3.909-10.829-9.057a6.001 6.001 0 0 1 .235-11.914C6.703 6.397 10.954 3 16 3c5.046 0 9.297 3.397 10.594 8.029a6 6 0 0 1 .235 11.914ZM6 13a4 4 0 0 0 0 8h1a9 9 0 1 0 18 0h1a4 4 0 0 0 0-8h-1.055a9 9 0 1 1-17.89 0H6Zm10 9a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/></g>`),
		g.Group(children),
	)
}