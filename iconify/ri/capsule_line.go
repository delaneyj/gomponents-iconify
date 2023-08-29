package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CapsuleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.779 4.222a6 6 0 0 1 0 8.485l-7.072 7.071a6 6 0 0 1-8.485-8.485l7.071-7.071a6 6 0 0 1 8.486 0Zm-5.657 11.313L8.466 9.878l-2.83 2.829a4 4 0 1 0 5.657 5.657l2.83-2.83Zm4.242-9.9a4 4 0 0 0-5.657 0L9.88 8.465l5.657 5.656l2.827-2.827a4 4 0 0 0 0-5.657Z"/>`),
		g.Group(children),
	)
}