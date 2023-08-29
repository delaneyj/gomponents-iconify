package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagStartBO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M50 0a3.5 3.5 0 0 0-3.5 3.5v80A3.5 3.5 0 0 0 50 87a3.5 3.5 0 0 0 3.5-3.5V47h43a3.5 3.5 0 0 0 3.5-3.5v-40A3.5 3.5 0 0 0 96.5 0h-46a3.5 3.5 0 0 0-.254.01A3.5 3.5 0 0 0 50 0zm4 7h39v33H54V7z" color="currentColor"/><path fill="currentColor" d="M38.004 76.792C27.41 78.29 20 81.872 20 87.143C20 94.243 32.381 100 50 100s30-5.756 30-12.857c0-5.272-7.41-8.853-18.003-10.35l-1.468 2.499C68.514 80.399 74 82.728 74 85.429c0 3.787-10.745 6.857-24 6.857s-24-3.07-24-6.857c-.001-2.692 5.45-5.018 13.459-6.13c-.484-.836-.97-1.67-1.455-2.507z" color="currentColor"/>`),
		g.Group(children),
	)
}