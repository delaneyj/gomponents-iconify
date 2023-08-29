package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaFlagOutline0"><g id="evaFlagOutline1"><path id="evaFlagOutline2" fill="currentColor" d="M19.27 4.68a1.79 1.79 0 0 0-1.6-.25a7.53 7.53 0 0 1-2.17.28a8.54 8.54 0 0 1-3.13-.78A10.15 10.15 0 0 0 8.5 3c-2.89 0-4 1-4.2 1.14a1 1 0 0 0-.3.72V20a1 1 0 0 0 2 0v-4.3a6.28 6.28 0 0 1 2.5-.41a8.54 8.54 0 0 1 3.13.78a10.15 10.15 0 0 0 3.87.93a7.66 7.66 0 0 0 3.5-.7a1.74 1.74 0 0 0 1-1.55V6.11a1.77 1.77 0 0 0-.73-1.43ZM18 14.59a6.32 6.32 0 0 1-2.5.41a8.36 8.36 0 0 1-3.13-.79a10.34 10.34 0 0 0-3.87-.92a9.51 9.51 0 0 0-2.5.29V5.42A6.13 6.13 0 0 1 8.5 5a8.36 8.36 0 0 1 3.13.79a10.34 10.34 0 0 0 3.87.92a9.41 9.41 0 0 0 2.5-.3Z"/></g></g>`),
		g.Group(children),
	)
}