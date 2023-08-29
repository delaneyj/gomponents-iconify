package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNotesSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M210.46 28.85a4 4 0 0 0-3.43-.73l-128 32A4 4 0 0 0 76 64v118.87A32 32 0 1 0 84 204V67.12l120-30v113.75a32 32 0 1 0 8 21.13V32a4 4 0 0 0-1.54-3.15ZM52 228a24 24 0 1 1 24-24a24 24 0 0 1-24 24Zm128-32a24 24 0 1 1 24-24a24 24 0 0 1-24 24Z"/>`),
		g.Group(children),
	)
}