package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 8v18h26V8zm4.313 2h17.375L16 15.781zM5 10.875l10.438 6.969l.562.343l.563-.343L27 10.875V24H5z"/>`),
		g.Group(children),
	)
}