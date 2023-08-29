package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropOfBlood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.64 2.548c-1.174-2.064-4.146-2.064-5.32 0l-6.84 12.03C5.578 16.168 5 18.006 5 20.012a10.983 10.983 0 0 0 12.908 10.816c4.026-.686 7.34-3.704 8.546-7.567c.992-3.175.435-6.094-.876-8.497l-.004-.008l-6.935-12.208Z"/>`),
		g.Group(children),
	)
}