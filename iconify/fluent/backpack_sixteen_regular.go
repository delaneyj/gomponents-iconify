package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackpackSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 6a2 2 0 0 0-2 2a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1a2 2 0 0 0-2-2H7ZM6 8a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1H6Zm2-7a2 2 0 0 0-2 2v.416A5.001 5.001 0 0 0 3 8v4.5A2.5 2.5 0 0 0 5.5 15h5a2.5 2.5 0 0 0 2.5-2.5V8a5.001 5.001 0 0 0-3-4.584V3a2 2 0 0 0-2-2Zm4 9H4V8a4 4 0 1 1 8 0v2Zm-6 2.5a.5.5 0 0 0 1 0V11h5v1.5a1.5 1.5 0 0 1-1.5 1.5h-5A1.5 1.5 0 0 1 4 12.5V11h2v1.5ZM8 3c-.342 0-.677.034-1 .1V3a1 1 0 0 1 2 0v.1A5.023 5.023 0 0 0 8 3Z"/>`),
		g.Group(children),
	)
}