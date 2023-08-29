package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="healthiconsForest0" d="M24 18c0 6.177-3.5 11.264-8 11.927V39h16v-8.954h-4.727a3.273 3.273 0 0 1-1.468-6.199a4.092 4.092 0 0 1 2.411-4.968a4.909 4.909 0 1 1 9.568 0a4.092 4.092 0 0 1 2.411 4.968a3.273 3.273 0 0 1-1.468 6.199H34V39h6.5a1.5 1.5 0 0 1 0 3h-33a1.5 1.5 0 0 1 0-3H14v-9.073C9.5 29.264 6 24.177 6 18c0-6.627 4.03-12 9-12s9 5.373 9 12Z"/></defs><g fill="currentColor"><use href="#healthiconsForest0"/><use href="#healthiconsForest0"/></g>`),
		g.Group(children),
	)
}