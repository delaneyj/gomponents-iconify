package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForUnitedArabEmirates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ed4c5c" d="M2 32c0 13.1 8.4 24.2 20 28.3V3.7C10.4 7.8 2 18.9 2 32z"/><path fill="#699635" d="M32 2c-3.5 0-6.9.6-10 1.7V22h38.3C56.2 10.4 45.1 2 32 2z"/><path fill="#f9f9f9" d="M60.3 22H22v20h38.3c1.1-3.1 1.7-6.5 1.7-10s-.6-6.9-1.7-10"/><path fill="#3e4347" d="M22 42v18.3c3.1 1.1 6.5 1.7 10 1.7c13.1 0 24.2-8.3 28.3-20H22z"/>`),
		g.Group(children),
	)
}