package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 8a4 4 0 0 0-4 4a3.91 3.91 0 0 0 .56 2H9.44a3.91 3.91 0 0 0 .56-2a4 4 0 1 0-4 4h12a4 4 0 0 0 0-8ZM6 14a2 2 0 1 1 2-2a2 2 0 0 1-2 2Zm12 0a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}