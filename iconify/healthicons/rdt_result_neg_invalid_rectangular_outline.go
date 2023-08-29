package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RdtResultNegInvalidRectangularOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6 19v10h36V19H6ZM4 29a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V19a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v10Z" clip-rule="evenodd"/><path d="M38 22a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm-5.5 3.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z"/><path fill-rule="evenodd" d="M26 23H11a1 1 0 1 0 0 2h15a1 1 0 1 0 0-2ZM8 24a3 3 0 0 0 3 3h15a3 3 0 1 0 0-6H11a3 3 0 0 0-3 3Zm3.593-12.974a1.207 1.207 0 0 0-.742.079a1.384 1.384 0 0 0-.609.537A1.621 1.621 0 0 0 10 12.5c0 .31.087.61.242.858c.156.248.37.432.609.537c.237.105.494.131.742.079c.248-.053.485-.185.677-.39a1 1 0 0 1 1.46 1.368a3.3 3.3 0 0 1-1.722.978a3.208 3.208 0 0 1-1.966-.206a3.383 3.383 0 0 1-1.495-1.304A3.62 3.62 0 0 1 8 12.5c0-.678.188-1.346.547-1.92a3.383 3.383 0 0 1 1.495-1.304a3.207 3.207 0 0 1 1.966-.206c.661.14 1.259.485 1.722.978a1 1 0 1 1-1.46 1.368a1.302 1.302 0 0 0-.677-.39ZM13 33a1 1 0 0 1 1-1h2.5a2.5 2.5 0 0 1 0 5H15v1a1 1 0 1 1-2 0v-5Zm2 2h1.5a.5.5 0 0 0 0-1H15v1Zm12-25a1 1 0 0 1 1-1h2.5a2.5 2.5 0 0 1 0 5H29v1a1 1 0 1 1-2 0v-5Zm2 2h1.5a.5.5 0 0 0 0-1H29v1Zm-6.5 27a1 1 0 0 0 .936-.649l1.5-4a1 1 0 1 0-1.872-.702l-.564 1.503l-.564-1.503a1 1 0 0 0-1.872.702l1.5 4A1 1 0 0 0 22.5 39Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M36 13v2a1 1 0 1 1-2 0v-2a1 1 0 1 1 0-2c0-.327.09-.794.389-1.21c.335-.465.885-.79 1.611-.79a1 1 0 0 1 0 2a1 1 0 0 1 0 2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}