package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 15a5 5 0 1 0-4 2h12a5 5 0 1 0-4-2h-4Zm-4 0a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm12 0a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}