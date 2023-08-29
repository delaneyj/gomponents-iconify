package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.75 12a1 1 0 0 0-.55-.89l-6.12-3.06v-4a3.08 3.08 0 1 0-6.16 0v4L2.8 11.11a1 1 0 0 0-.55.89v3.33a1 1 0 0 0 .43.83a1 1 0 0 0 .92.11l5.32-2V18l-1.82.6a1 1 0 0 0-.68.95V22a1 1 0 0 0 .3.71a1 1 0 0 0 .7.29h9.17a1 1 0 0 0 1-1v-2.5a1 1 0 0 0-.68-.95L15.08 18v-3.72l5.32 2a1 1 0 0 0 .92-.11a1 1 0 0 0 .43-.83Zm-7.31-.1a1 1 0 0 0-.93.11a1 1 0 0 0-.43.82v5.84a1 1 0 0 0 .69.95l1.81.6V21H8.41v-.78l1.81-.6a1 1 0 0 0 .69-.95v-5.84a1 1 0 0 0-.43-.82a1 1 0 0 0-.93-.11l-5.31 2v-1.28l6.11-3.06a1 1 0 0 0 .56-.89V4.08a1.08 1.08 0 1 1 2.16 0v4.59a1 1 0 0 0 .56.89l6.11 3.06v1.27Z"/>`),
		g.Group(children),
	)
}