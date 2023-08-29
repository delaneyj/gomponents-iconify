package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RdtResultNegNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsRdtResultNegNegative0)"><path d="M38 22a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm-7 2a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Z"/><path fill-rule="evenodd" d="M26 21a3 3 0 1 1 0 6H11a3 3 0 1 1 0-6h15Zm-13 2h-2a1 1 0 1 0 0 2h2v-2Zm2 2h11a1 1 0 1 0 0-2H15v2Z" clip-rule="evenodd"/><path d="M15 35h1.5a.5.5 0 0 0 0-1H15v1Zm15.5-23H29v-1h1.5a.5.5 0 0 1 0 1Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm10 30a6 6 0 0 1 0-12h28a6 6 0 0 1 0 12H10Zm.85-18.895c.238-.105.495-.131.743-.079c.248.053.485.185.677.39a1 1 0 0 0 1.46-1.368a3.301 3.301 0 0 0-1.722-.978a3.207 3.207 0 0 0-1.966.206a3.383 3.383 0 0 0-1.495 1.304A3.62 3.62 0 0 0 8 12.5c0 .678.188 1.346.547 1.92c.36.574.877 1.031 1.495 1.304a3.208 3.208 0 0 0 1.966.206a3.3 3.3 0 0 0 1.722-.978a1 1 0 1 0-1.46-1.368c-.192.205-.43.337-.677.39a1.207 1.207 0 0 1-.742-.079a1.384 1.384 0 0 1-.609-.537A1.62 1.62 0 0 1 10 12.5c0-.31.087-.61.242-.858c.156-.248.37-.432.609-.537ZM13 33a1 1 0 0 1 1-1h2.5a2.5 2.5 0 0 1 0 5H15v1a1 1 0 1 1-2 0v-5ZM28 9a1 1 0 0 0-1 1v5a1 1 0 1 0 2 0v-1h1.5a2.5 2.5 0 0 0 0-5H28Zm-5.5 30a1 1 0 0 0 .936-.649l1.5-4a1 1 0 1 0-1.872-.702l-.564 1.503l-.564-1.503a1 1 0 0 0-1.872.702l1.5 4A1 1 0 0 0 22.5 39ZM35 16a1 1 0 0 0 1-1v-2a1 1 0 0 0 0-2a1 1 0 0 0 0-2c-.726 0-1.276.325-1.611.79A2.116 2.116 0 0 0 34 11a1 1 0 1 0 0 2v2a1 1 0 0 0 1 1Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsRdtResultNegNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}