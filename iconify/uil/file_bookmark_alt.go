package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileBookmarkAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 10h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2Zm4 6h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Zm0-4h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Zm6.92-2.62a1 1 0 0 0-.21-1.09l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0l-.28-.1H5.5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h6a1 1 0 0 0 0-2h-6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h4a1 1 0 0 0 .92-.62ZM13.5 8a1 1 0 0 1-1-1V5.41L15.09 8Zm7 4h-5a1 1 0 0 0-1 1v8a1 1 0 0 0 .53.88a1 1 0 0 0 1-.05l1.97-1.3l2 1.3a1 1 0 0 0 1.5-.83v-8a1 1 0 0 0-1-1Zm-1 7.12l-.94-.63a1 1 0 0 0-1.12 0l-.94.63V14h3Z"/>`),
		g.Group(children),
	)
}