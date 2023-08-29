package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BicycleShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10 1a1 1 0 1 0 0 2a1 1 0 0 0 0-2zM8.145 2.994a.5.5 0 0 0-.348.143l-2.64 2.5a.5.5 0 0 0 .042.763L7 7.75v2.75c-.01.676 1.01.676 1 0v-3a.5.5 0 0 0-.2-.4l-.767-.577l1.818-1.72l.749.998A.5.5 0 0 0 10 6h1.5c.676.01.676-1.01 0-1h-1.25L9.5 4l-.6-.8a.5.5 0 0 0-.384-.206h-.371zM3 7a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm9 0a3 3 0 1 0 0 6a3 3 0 0 0 0-6zM3 8a2 2 0 1 1 0 4a2 2 0 0 1 0-4zm9 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4z"/>`),
		g.Group(children),
	)
}