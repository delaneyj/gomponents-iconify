package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nuclear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path fill="currentColor" d="M11.1 17a4.996 4.996 0 0 0 1.585 2.743l-2.024 3.504A8.991 8.991 0 0 1 7.056 17H11.1Zm13.845 0a8.992 8.992 0 0 1-3.607 6.247l-2.023-3.504A4.996 4.996 0 0 0 20.9 17h4.045ZM16 11c-.553 0-1.086.09-1.584.256l-2.023-3.504A8.966 8.966 0 0 1 16 7c1.284 0 2.504.268 3.607.752l-2.023 3.504A4.994 4.994 0 0 0 16 11Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 16a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}