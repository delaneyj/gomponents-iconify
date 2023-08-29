package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BandaidsThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m178.91 128l30.55-30.54a36 36 0 0 0-50.92-50.92L128 77.09L97.46 46.54a36 36 0 0 0-50.92 50.92L77.09 128l-30.55 30.54a36 36 0 0 0 50.92 50.92L128 178.91l30.54 30.55a36 36 0 0 0 50.92-50.92ZM164.2 52.2a28 28 0 0 1 39.6 39.6l-30.55 30.54l-39.59-39.59Zm3.4 75.8L128 167.6L88.4 128L128 88.4ZM52.2 91.8a28 28 0 0 1 39.6-39.6l30.54 30.55l-39.59 39.59Zm39.6 112a28 28 0 0 1-39.6-39.6l30.55-30.54l39.59 39.59Zm112 0a28 28 0 0 1-39.6 0l-30.54-30.55l39.59-39.59l30.55 30.54a28 28 0 0 1 0 39.6ZM120 128a8 8 0 1 1 8 8a8 8 0 0 1-8-8Z"/>`),
		g.Group(children),
	)
}