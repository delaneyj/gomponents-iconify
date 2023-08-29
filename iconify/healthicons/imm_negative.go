package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImmNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsImmNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0Zm-6 24c0 9.941-8.059 18-18 18S6 33.941 6 24S14.059 6 24 6s18 8.059 18 18Zm2 0c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4s20 8.954 20 20Zm-20 7a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm0 2a9 9 0 1 0 0-18a9 9 0 0 0 0 18Zm-1.5-4a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm-8-18a3.5 3.5 0 1 1 0 7a3.5 3.5 0 0 1 0-7Zm21 24a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7ZM31 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6ZM14 32a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsImmNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}