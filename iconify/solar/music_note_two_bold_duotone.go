package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9 13.25a4.75 4.75 0 1 0 0 9.5a4.75 4.75 0 0 0 0-9.5Z"/><path fill-rule="evenodd" d="M13 1.25a.75.75 0 0 1 .75.75c0 2.9 2.35 5.25 5.25 5.25a.75.75 0 0 1 0 1.5A6.75 6.75 0 0 1 12.25 2a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/><path d="M12.25 14.536V2c0 1.607.562 3.084 1.5 4.243V18a4.737 4.737 0 0 0-1.5-3.464Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}