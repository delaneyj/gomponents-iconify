package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469.3 469.3H42.7V42.7H256L298.7 0h-256C19.1 0 0 19.1 0 42.7v426.7C0 492.9 19.1 512 42.7 512h426.7c23.6 0 42.7-19.1 42.7-42.7V320l-42.7 42.7v106.6zm-384-42.6C149.1 255.7 234.7 256 362.7 256v85.3L512 192L362.7 42.7V128C85.3 128 85.1 341.1 85.3 426.7z"/>`),
		g.Group(children),
	)
}