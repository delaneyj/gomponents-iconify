package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileImport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 20H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v2a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.32 1.32 0 0 0-.19-.29l-6-6a1.32 1.32 0 0 0-.29-.19a.32.32 0 0 0-.09 0l-.31-.1H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h5a1 1 0 0 0 0-2Zm2-14.59L15.59 8H14a1 1 0 0 1-1-1ZM19 15h-5.59l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3a1.15 1.15 0 0 0-.21.33a1 1 0 0 0 0 .76a.93.93 0 0 0 .21.33l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L13.41 17H19a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}