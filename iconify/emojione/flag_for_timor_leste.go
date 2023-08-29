package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTimorLeste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ed4c5c" d="M32 2c-7.6 0-14.6 2.9-19.9 7.6v44.9C17.4 59.1 24.4 62 32 62c16.6 0 30-13.4 30-30S48.6 2 32 2"/><path fill="#ffce31" d="M12.1 9.6c-1.3 1.2-2.5 2.5-3.6 3.9v37.1c1.1 1.4 2.3 2.7 3.6 3.9L42 32L12.1 9.6z"/><path fill="#3e4347" d="M8.4 13.4C4.4 18.5 2 25 2 32s2.4 13.5 6.4 18.6L27 32L8.4 13.4z"/><path fill="#fff" d="m10.3 33.9l2.4 5.1l1-5.3l5.3-.4l-4.8-3l1-5.3l-3.9 3.5l-4.8-2.9l2.4 5.1L5 34.3z"/>`),
		g.Group(children),
	)
}