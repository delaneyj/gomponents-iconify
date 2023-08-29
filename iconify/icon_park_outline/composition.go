package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Composition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="m39.828 21.786l-17.86 17.86c-3.414 3.414-9.03 3.332-12.545-.182c-3.515-3.515-3.597-9.132-.182-12.546L27.1 9.058" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="m39.828 21.786l-17.86 17.86c-3.414 3.414-9.03 3.332-12.545-.182c-3.515-3.515-3.597-9.132-.182-12.546L27.1 9.058M8.734 27.424l26.144-.688"/><path fill="currentColor" d="M29.393 20.907a2 2 0 1 0-2.828-2.829a2 2 0 0 0 2.828 2.829Zm-5.657 2.828a2 2 0 1 0-2.828-2.828a2 2 0 0 0 2.828 2.828Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="m23.564 5.522l19.8 19.8"/></g>`),
		g.Group(children),
	)
}