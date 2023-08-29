package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.75c-3.241 0-5.756 2.03-6.185 4.5H8a.75.75 0 0 1 .75.75v7a.75.75 0 0 1-.75.75H5A2.75 2.75 0 0 1 2.25 14v-3a2.75 2.75 0 0 1 2.035-2.656C4.667 4.84 8.074 2.25 12 2.25c3.926 0 7.333 2.59 7.715 6.094A2.751 2.751 0 0 1 21.75 11v3a2.751 2.751 0 0 1-2.045 2.659A4.751 4.751 0 0 1 15 20.75h-1.145a2 2 0 1 1 0-1.5H15a3.251 3.251 0 0 0 3.163-2.5H16a.75.75 0 0 1-.75-.75V9a.75.75 0 0 1 .75-.75h2.185c-.429-2.47-2.944-4.5-6.185-4.5Zm-7 6c-.69 0-1.25.56-1.25 1.25v3c0 .69.56 1.25 1.25 1.25h2.25v-5.5H5ZM20.25 11c0-.69-.56-1.25-1.25-1.25h-2.25v5.5H19c.69 0 1.25-.56 1.25-1.25v-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}