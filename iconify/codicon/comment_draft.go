package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentDraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 2.001H13v1h1v1h1v-1.5l-.5-.5Zm-5.5 0h2v1H9v-1Zm-4 0h2v1H5v-1Zm9 8v2h.5l.5-.5v-1.5h-1Zm-2 2v-1h-2v1h2Zm-4-1h-.5l-.354.146L5 13.294v-1.793l-.5-.5H4v3.5l.854.354l2.853-2.854H8v-1Zm7-3v-2h-1v2h1Zm-13 3v-1H1v1.5l.5.5H2v-1Zm0-3v-2H1v2h1Zm0-5v1H1v-1.5l.5-.5H3v1H2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}