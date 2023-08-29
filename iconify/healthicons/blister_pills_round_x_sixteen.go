package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlisterPillsRoundXSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 4a4 4 0 0 0-4 4v15a1 1 0 1 1 0 2v15a4 4 0 0 0 4 4h16a4 4 0 0 0 4-4V25a1 1 0 1 1 0-2V8a4 4 0 0 0-4-4H16Zm0 20a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM18 10.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm3.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm6.5-1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm3.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM18 15.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm3.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm6.5-1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm3.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM18 31.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm3.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm6.5-1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm3.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM18 36.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm3.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm6.5-1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm3.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}