package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NineNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthicons9Negative0)"><path d="M25.968 22C28.205 22 30 20.2 30 18s-1.795-4-4.032-4h-3.936C19.795 14 18 15.8 18 18s1.795 4 4.032 4h3.936Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM18.228 31.33A4.033 4.033 0 0 0 22.032 34h3.936C28.205 34 30 32.2 30 30v-5.08A8.022 8.022 0 0 1 25.968 26h-3.936C17.606 26 14 22.428 14 18s3.606-8 8.032-8h3.936C30.394 10 34 13.572 34 18v12c0 4.428-3.606 8-8.032 8h-3.936a8.032 8.032 0 0 1-7.573-5.33a2 2 0 0 1 3.769-1.34Z" clip-rule="evenodd"/></g><defs><clipPath id="healthicons9Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}