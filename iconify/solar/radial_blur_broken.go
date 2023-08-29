package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadialBlurBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" d="M15.5 14.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm0-4.5a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm-4.5 4.5a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm0-4.5a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/><path fill="currentColor" d="M15 18.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm-4.5 0a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/><path fill="currentColor" d="M15 18.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm0-12.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm-4.5 12.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm0-12.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0ZM18.25 9a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5ZM5.75 9a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm12.5 4.5a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm-12.5 0a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7 3.338A9.954 9.954 0 0 1 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12c0-1.821.487-3.53 1.338-5"/></g>`),
		g.Group(children),
	)
}