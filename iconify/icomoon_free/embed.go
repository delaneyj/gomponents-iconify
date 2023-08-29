package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Embed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m9 11.5l1.5 1.5l5-5l-5-5L9 4.5L12.5 8zm-2-7L5.5 3l-5 5l5 5L7 11.5L3.5 8z"/>`),
		g.Group(children),
	)
}