package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M28.6 17.5h6.9l10.3 29h-6.6l-1.9-6H26.6l-2 6h-6.3l10.3-29m-.4 18h7.4L32 24.1l-3.8 11.4"/>`),
		g.Group(children),
	)
}