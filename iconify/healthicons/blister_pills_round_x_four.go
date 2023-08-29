package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlisterPillsRoundXFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 4a4 4 0 0 0-4 4v15a1 1 0 1 1 0 2v15a4 4 0 0 0 4 4h16a4 4 0 0 0 4-4V25a1 1 0 1 1 0-2V8a4 4 0 0 0-4-4H16Zm0 20a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM22 14a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm7 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm-7 17a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm7 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}