package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaEject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M97.088 286.56h317.824a24 24 0 0 0 18.432-39.371L274.433 56.63a24 24 0 0 0-36.866 0L78.656 247.189a24 24 0 0 0 18.432 39.371ZM256 84.491L397.824 254.56H114.176ZM416 328H96a24.027 24.027 0 0 0-24 24v64a24.027 24.027 0 0 0 24 24h320a24.027 24.027 0 0 0 24-24v-64a24.027 24.027 0 0 0-24-24Zm-8 80H104v-48h304Z"/>`),
		g.Group(children),
	)
}