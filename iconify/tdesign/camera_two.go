package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3.5a7 7 0 1 0 0 14a7 7 0 0 0 0-14Zm-9 7a9 9 0 1 1 10 8.945v1.01a9.937 9.937 0 0 0 4.626-1.682l.826-.563l1.127 1.652l-.826.564A11.937 11.937 0 0 1 13 22.464v.036h-.652a12.216 12.216 0 0 1-.696 0H11v-.036a11.937 11.937 0 0 1-5.753-2.038l-.826-.564l1.127-1.652l.826.563A9.937 9.937 0 0 0 11 20.456v-1.01c-4.5-.498-8-4.313-8-8.946Zm9-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0Z"/>`),
		g.Group(children),
	)
}