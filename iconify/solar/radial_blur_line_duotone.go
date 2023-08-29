package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadialBlurLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M3.34 17c2.761 4.783 8.877 6.421 13.66 3.66a9.956 9.956 0 0 0 4.197-4.731a9.985 9.985 0 0 0-.537-8.93a9.985 9.985 0 0 0-7.464-4.928A9.956 9.956 0 0 0 7 3.339C2.217 6.101.58 12.217 3.34 17Z" opacity=".5"/><path stroke="currentColor" d="M15.5 14.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm0-4.5a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm-4.5 4.5a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm0-4.5a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/><path fill="currentColor" d="M15 18.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm-4.5 0a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/><path fill="currentColor" d="M15 18.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm0-12.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm-4.5 12.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm0-12.5a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0ZM18.25 9a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5ZM5.75 9a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm12.5 4.5a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm-12.5 0a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Z"/></g>`),
		g.Group(children),
	)
}