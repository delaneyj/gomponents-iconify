package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthicons6Negative0)"><path d="M22.032 26C19.795 26 18 27.8 18 30s1.795 4 4.032 4h3.936C28.205 34 30 32.2 30 30s-1.795-4-4.032-4h-3.936Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM29.772 16.67A4.033 4.033 0 0 0 25.968 14h-3.936C19.795 14 18 15.8 18 18v5.08A8.022 8.022 0 0 1 22.032 22h3.936C30.394 22 34 25.572 34 30s-3.606 8-8.032 8h-3.936C17.606 38 14 34.428 14 30V18c0-4.428 3.606-8 8.032-8h3.936a8.032 8.032 0 0 1 7.573 5.33a2 2 0 0 1-3.769 1.34Z" clip-rule="evenodd"/></g><defs><clipPath id="healthicons6Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}