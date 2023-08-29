package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisinfectingWipesNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsDisinfectingWipesNegative0)"><path fill-rule="evenodd" d="M24.581 6.814c1.422-1.015 3.337-1.063 5.062-.766c.807.14 1.526.346 2.055.525L27.21 17.9h-6.36L17.753 7.804l.03.01c1.657.603 3.972.7 6.732-.956a.988.988 0 0 0 .066-.044Zm3.842 3.07a1 1 0 1 0-1.846-.769l-2.5 6a1 1 0 1 0 1.846.77l2.5-6Z" clip-rule="evenodd"/><path d="M20 29c3.031.63 4.795.61 8 0v6.165c-3.119 1-4.906.9-8 0V29Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM33.93 6.368a1 1 0 0 0-.516-1.278l-.003-.002l-.006-.002l-.016-.007a6.595 6.595 0 0 0-.258-.109a15.967 15.967 0 0 0-3.149-.893c-1.89-.326-4.456-.375-6.528 1.085c-2.225 1.323-3.9 1.169-4.987.773a4.783 4.783 0 0 1-1.304-.722a3.925 3.925 0 0 1-.396-.356l-.012-.013a1 1 0 0 0-1.71.95l3.217 10.487C17.467 16.768 17 17.361 17 18c0 1.657 3.134 3 7 3s7-1.343 7-3c0-.58-.385-1.123-1.052-1.582l3.982-10.05ZM31 21.571c-1.804.884-4.274 1.43-7 1.43s-5.196-.546-7-1.43V41c0 1.657 3.134 3 7 3s7-1.343 7-3V21.57Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsDisinfectingWipesNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}