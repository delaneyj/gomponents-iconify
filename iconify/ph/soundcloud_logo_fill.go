package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundcloudLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M24 120v48a8 8 0 0 1-16 0v-48a8 8 0 0 1 16 0Zm24-32a8 8 0 0 0-8 8v96a8 8 0 0 0 16 0V96a8 8 0 0 0-8-8Zm32-8a8 8 0 0 0-8 8v104a8 8 0 0 0 16 0V88a8 8 0 0 0-8-8Zm32-32a8 8 0 0 0-8 8v136a8 8 0 0 0 16 0V56a8 8 0 0 0-8-8Zm110.84 58.34A80 80 0 0 0 144 40a8 8 0 0 0-8 8v144a8 8 0 0 0 8 8h64a48 48 0 0 0 14.84-93.66Z"/>`),
		g.Group(children),
	)
}