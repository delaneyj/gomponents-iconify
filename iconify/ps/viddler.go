package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viddler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M351 221q46 0 78.5-32.5T462 110t-32.5-78T351 0q-42 0-73 27t-37 67h-3q-6-40-37.5-67T128 0Q81 0 48.5 32T16 110t32 78t78 33h3v37L2 224v147l127-38v51h188l-55-163h89zm-223-80q-14 0-23-9t-9-22t9.5-22t22.5-9t22 9t9 22t-9 22t-22 9zm194-31q0-13 9-22t22-9t22.5 9t9.5 22t-9 22t-23 9q-13 0-22-9t-9-22z"/>`),
		g.Group(children),
	)
}