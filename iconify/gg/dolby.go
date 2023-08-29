package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dolby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 4v16h24V4H0Zm10 8a4 4 0 0 0-4-4H4v8h2a4 4 0 0 0 4-4Zm8 4h2V8h-2a4 4 0 0 0 0 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}