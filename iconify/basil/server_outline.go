package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 7.5a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm3 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm-2 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/><path fill="currentColor" fill-rule="evenodd" d="M5.947 3.25A2.75 2.75 0 0 0 3.197 6v3c0 .788.332 1.499.863 2a2.742 2.742 0 0 0-.863 2v3a2.75 2.75 0 0 0 2.75 2.75h5.303v1.5H5a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 0-1.5h-6.25v-1.5h5.197a2.75 2.75 0 0 0 2.75-2.75v-3c0-.788-.331-1.499-.862-2a2.743 2.743 0 0 0 .862-2V6a2.75 2.75 0 0 0-2.75-2.75h-12Zm0 7h12c.69 0 1.25-.56 1.25-1.25V6c0-.69-.56-1.25-1.25-1.25h-12c-.69 0-1.25.56-1.25 1.25v3c0 .69.56 1.25 1.25 1.25Zm0 1.5c-.69 0-1.25.56-1.25 1.25v3c0 .69.56 1.25 1.25 1.25h12c.69 0 1.25-.56 1.25-1.25v-3c0-.69-.56-1.25-1.25-1.25h-12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}