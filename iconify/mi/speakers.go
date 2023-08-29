package mi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speakers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4zm14 0H6v16h12V4z"/><path d="M12 12a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0zm5.5-6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0z"/></g>`),
		g.Group(children),
	)
}