package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M286.903 256L416 177.877v-36.762l-.283-.469L272 227.617V72h-32v155.617L96.283 140.646l-.283.469v36.762L225.097 256L96 334.123v36.762l.283.469L240 284.383V440h32V284.383l143.717 86.971l.283-.469v-36.762L286.903 256z"/>`),
		g.Group(children),
	)
}