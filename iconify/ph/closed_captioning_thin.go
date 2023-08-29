package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptioningThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 52H32a12 12 0 0 0-12 12v128a12 12 0 0 0 12 12h192a12 12 0 0 0 12-12V64a12 12 0 0 0-12-12Zm4 140a4 4 0 0 1-4 4H32a4 4 0 0 1-4-4V64a4 4 0 0 1 4-4h192a4 4 0 0 1 4 4Zm-112.54-38.29a4 4 0 0 1-1.46 5.47a36 36 0 1 1 0-62.36a4 4 0 0 1-4 6.92a28 28 0 1 0 0 48.52a4 4 0 0 1 5.46 1.45Zm80 0a4 4 0 0 1-1.46 5.47a36 36 0 1 1 0-62.36a4 4 0 0 1-4 6.92a28 28 0 1 0 0 48.52a4 4 0 0 1 5.46 1.45Z"/>`),
		g.Group(children),
	)
}