package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feSpeaker0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSpeaker1" fill="currentColor"><path id="feSpeaker2" d="M18 20a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v16Zm-6-8a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`),
		g.Group(children),
	)
}