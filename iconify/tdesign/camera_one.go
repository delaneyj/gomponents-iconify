package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a7 7 0 1 0 0 14a7 7 0 0 0 0-14Zm-9 7a9 9 0 1 1 13.148 7.989L16.805 21H19v2H5v-2h2.195l.657-3.011A8.999 8.999 0 0 1 3 10Zm6.74 8.714L9.243 21h5.516l-.499-2.286A9.012 9.012 0 0 1 12 19c-.78 0-1.537-.1-2.26-.286ZM12 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0Z"/>`),
		g.Group(children),
	)
}