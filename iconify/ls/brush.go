package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m680 28l11 12s54 54 5 104c-42 41-89 43-145 79c-67 41-103 100-95 135c15 65 92 118 25 186c-5 4-15 16-19 20c0 0 3 21-20 21c0 0 5 22-20 22L270 455l-5-5l-151-151c0-24 21-21 21-21c0-16 14-19 20-19c0 0 14-15 20-21c67-68 122 11 185 25c36 9 96-27 137-94c35-57 36-104 78-146c50-50 105 5 105 5zm-86 29c-19 19-19 51 0 70s52 20 71 1s19-53 0-72s-52-18-71 1zM3 453l-3-40l82-82l305 307l-81 83l-41-5l-4-48l-48-4l-3-51l-49-2l-4-49l-48-3l-5-49l-49-4l-3-49z"/>`),
		g.Group(children),
	)
}