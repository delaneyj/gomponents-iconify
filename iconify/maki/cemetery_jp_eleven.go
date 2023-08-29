package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CemeteryJpEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8.205 7.905h-2v-5.6a.736.736 0 0 0-.8-.8a.713.713 0 0 0-.794.62a.734.734 0 0 0-.006.08v5.7h-1.8a.736.736 0 0 0-.8.8a.714.714 0 0 0 .623.795a.701.701 0 0 0 .077.005h5.5a.736.736 0 0 0 .8-.666a.75.75 0 0 0 0-.134a.768.768 0 0 0-.735-.8h-.065z" fill="currentColor"/>`),
		g.Group(children),
	)
}