package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelDownAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.5 4.5v3h15v15.355l-3.781-3.78l-2.145 2.144l6.358 6.357L21 28.598l1.068-1.022l6.358-6.357l-2.145-2.145l-3.781 3.781V4.5h-18zm1 1h16v19.77l4.781-4.782l.73.73l-5.642 5.643l-.369.354l-.37-.354l-5.642-5.642l.73-.73l4.782 4.78V6.5h-15v-1z"/>`),
		g.Group(children),
	)
}