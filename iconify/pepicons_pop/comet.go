package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.29 13.093a1 1 0 1 0 1.52-1.3a1 1 0 0 0-1.52 1.3ZM9 14.723a3 3 0 1 1-3.9-4.56a3 3 0 0 1 3.9 4.56Z"/><path d="M10.907 16.949A5.921 5.921 0 1 1 3.23 7.934l7.08-5.836c1.078-.89 2.68.05 2.43 1.425l-.263 1.445l4.558-3.799c1.304-1.086 3.125.487 2.24 1.934l-2.858 4.678l.896-.1c1.458-.162 2.258 1.649 1.156 2.618l-7.562 6.65Zm-6.89-1.897a3.92 3.92 0 0 0 5.57.395l6.37-5.603l-.333.037c-1.24.138-2.096-1.209-1.445-2.273l1.812-2.965l-3.292 2.743c-1.077.898-2.687-.041-2.436-1.42l.264-1.456l-6.024 4.967a3.921 3.921 0 0 0-.485 5.575Z"/></g>`),
		g.Group(children),
	)
}