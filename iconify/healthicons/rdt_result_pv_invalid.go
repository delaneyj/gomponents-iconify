package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RdtResultPvInvalid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M28 9a1 1 0 0 0-1 1v5a1 1 0 1 0 2 0v-1h1.5a2.5 2.5 0 0 0 0-5H28Zm2.5 3H29v-1h1.5a.5.5 0 0 1 0 1Z" clip-rule="evenodd"/><path d="M11.593 11.026a1.207 1.207 0 0 0-.742.079a1.384 1.384 0 0 0-.609.537A1.621 1.621 0 0 0 10 12.5c0 .31.087.61.242.858c.156.248.37.432.609.537c.237.105.494.131.742.079c.248-.053.485-.185.677-.39a1 1 0 0 1 1.46 1.368a3.3 3.3 0 0 1-1.722.978a3.208 3.208 0 0 1-1.966-.206a3.383 3.383 0 0 1-1.495-1.304A3.62 3.62 0 0 1 8 12.5c0-.678.188-1.346.547-1.92a3.383 3.383 0 0 1 1.495-1.304a3.207 3.207 0 0 1 1.966-.206c.661.14 1.259.485 1.722.978a1 1 0 1 1-1.46 1.368a1.302 1.302 0 0 0-.677-.39Z"/><path fill-rule="evenodd" d="M14 32a1 1 0 0 0-1 1v5a1 1 0 1 0 2 0v-1h1.5a2.5 2.5 0 0 0 0-5H14Zm2.5 3H15v-1h1.5a.5.5 0 0 1 0 1Z" clip-rule="evenodd"/><path d="M23.436 38.351a1 1 0 0 1-1.872 0l-1.5-4a1 1 0 1 1 1.872-.702l.564 1.503l.564-1.503a1 1 0 1 1 1.872.702l-1.5 4ZM35 16a1 1 0 0 0 1-1v-2a1 1 0 0 0 0-2a1 1 0 0 0 0-2c-.726 0-1.276.325-1.611.79A2.116 2.116 0 0 0 34 11a1 1 0 1 0 0 2v2a1 1 0 0 0 1 1Zm-22 7h-2a1 1 0 1 0 0 2h2v-2Zm2 0v2h7v-2h-7Zm9 2h2a1 1 0 1 0 0-2h-2v2Z"/><path fill-rule="evenodd" d="M10 30a6 6 0 0 1 0-12h28a6 6 0 0 1 0 12H10Zm30-6a2 2 0 1 0-4 0a2 2 0 0 0 4 0Zm-7.5 1.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM26 21a3 3 0 1 1 0 6H11a3 3 0 1 1 0-6h15Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}