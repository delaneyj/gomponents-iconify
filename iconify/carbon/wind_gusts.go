package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindGusts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.316 8.051l-18-6a1 1 0 0 0-.916.149L4 7V2H2v28h2V11l6.4 4.8a1 1 0 0 0 .916.149l18-6a1 1 0 0 0 0-1.897ZM10 13L4.667 9L10 5Zm4-.054l-2 .667V4.387l2 .667Zm4-1.333l-2 .666V5.721l2 .666Zm2-.667V7.054L25.838 9Z"/><path fill="currentColor" d="M20 22a4 4 0 0 0-8 0h2a2 2 0 1 1 2 2H8v2h8a4.005 4.005 0 0 0 4-4Z"/><path fill="currentColor" d="M26 22a4.005 4.005 0 0 0-4 4h2a2 2 0 1 1 2 2H12v2h14a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}