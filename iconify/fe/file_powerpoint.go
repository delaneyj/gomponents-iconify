package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePowerpoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFilePowerpoint0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFilePowerpoint1" fill="currentColor"><path id="feFilePowerpoint2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4ZM10 16v2H8v-8h5a3 3 0 0 1 0 6h-3Zm0-4v2h3a1 1 0 0 0 0-2h-3Z"/></g></g>`),
		g.Group(children),
	)
}