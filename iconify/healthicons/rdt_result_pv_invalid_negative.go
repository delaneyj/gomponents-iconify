package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RdtResultPvInvalidNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsRdtResultPvInvalidNegative0)"><path d="M38 22a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm-7 2a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Z"/><path fill-rule="evenodd" d="M29 24a3 3 0 0 0-3-3H11a3 3 0 1 0 0 6h15a3 3 0 0 0 3-3Zm-18-1h2v2h-2a1 1 0 1 1 0-2Zm4 2v-2h7v2h-7Zm11 0h-2v-2h2a1 1 0 1 1 0 2Z" clip-rule="evenodd"/><path d="M16.5 35H15v-1h1.5a.5.5 0 0 1 0 1Zm14-23H29v-1h1.5a.5.5 0 0 1 0 1Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM4 24a6 6 0 0 0 6 6h28a6 6 0 0 0 0-12H10a6 6 0 0 0-6 6Zm7.593-12.974a1.207 1.207 0 0 0-.742.079a1.384 1.384 0 0 0-.609.537A1.62 1.62 0 0 0 10 12.5c0 .31.087.61.242.858c.156.248.37.432.609.537c.237.105.494.131.742.079c.248-.053.485-.185.677-.39a1 1 0 0 1 1.46 1.368a3.3 3.3 0 0 1-1.722.978a3.208 3.208 0 0 1-1.966-.206a3.383 3.383 0 0 1-1.495-1.304A3.62 3.62 0 0 1 8 12.5c0-.678.188-1.346.547-1.92a3.383 3.383 0 0 1 1.495-1.304a3.207 3.207 0 0 1 1.966-.206c.661.14 1.259.485 1.722.978a1 1 0 1 1-1.46 1.368a1.302 1.302 0 0 0-.677-.39ZM14 32a1 1 0 0 0-1 1v5a1 1 0 1 0 2 0v-1h1.5a2.5 2.5 0 0 0 0-5H14Zm13-22a1 1 0 0 1 1-1h2.5a2.5 2.5 0 0 1 0 5H29v1a1 1 0 1 1-2 0v-5Zm-3.564 28.351a1 1 0 0 1-1.872 0l-1.5-4a1 1 0 0 1 1.872-.702l.564 1.503l.564-1.503a1 1 0 1 1 1.872.702l-1.5 4ZM36 15a1 1 0 1 1-2 0v-2a1 1 0 1 1 0-2c0-.327.09-.794.389-1.21c.335-.465.885-.79 1.611-.79a1 1 0 0 1 0 2a1 1 0 0 1 0 2v2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsRdtResultPvInvalidNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}