package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VocabularyOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 3h3a2 2 0 0 1 2 2a2 2 0 0 1 2-2h6a1 1 0 0 1 1 1v13m-2 2h-5a2 2 0 0 0-2 2a2 2 0 0 0-2-2H4a1 1 0 0 1-1-1V4c0-.279.114-.53.298-.712M12 5v3m0 4v9M7 11h1m8-4h1m-1 4h1M3 3l18 18"/>`),
		g.Group(children),
	)
}