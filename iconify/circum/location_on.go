package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.933a1.715 1.715 0 0 1-1.384-.691L5.555 14.5a7.894 7.894 0 1 1 12.885-.009l-5.055 6.749a1.717 1.717 0 0 1-1.385.693Zm-.008-18.867a6.81 6.81 0 0 0-4.578 1.749a6.891 6.891 0 0 0-1.05 9.1l5.051 6.727a.725.725 0 0 0 .584.292a.732.732 0 0 0 .586-.292l5.044-6.734A6.874 6.874 0 0 0 12.81 3.113a7.277 7.277 0 0 0-.818-.047Z"/><path fill="currentColor" d="M12 12.5a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5Zm0-4a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 0 0 12 8.5Z"/>`),
		g.Group(children),
	)
}