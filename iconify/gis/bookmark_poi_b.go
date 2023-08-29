package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkPoiB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M30.027.027a1.5 1.5 0 0 0-1.5 1.5V2h-.496a2.5 2.5 0 0 0-2.5 2.5a2.5 2.5 0 0 0 2.5 2.5h.496v49.527a1.5 1.5 0 0 0 2.383 1.213L46.5 46.395v36.773a3.5 3.5 0 1 0 7 0V46.395L69.09 57.74a1.5 1.5 0 0 0 2.383-1.213V7h.523a2.5 2.5 0 0 0 2.5-2.5a2.5 2.5 0 0 0-2.5-2.5h-.523v-.473a1.5 1.5 0 0 0-1.5-1.5H30.027z" color="currentColor"/><path fill="currentColor" d="M38.004 76.792C27.41 78.29 20 81.872 20 87.143C20 94.243 32.381 100 50 100s30-5.756 30-12.857c0-5.272-7.41-8.853-18.003-10.35l-1.468 2.499C68.514 80.399 74 82.728 74 85.429c0 3.787-10.745 6.857-24 6.857s-24-3.07-24-6.857c-.001-2.692 5.45-5.018 13.459-6.13c-.484-.836-.97-1.67-1.455-2.507z" color="currentColor"/>`),
		g.Group(children),
	)
}