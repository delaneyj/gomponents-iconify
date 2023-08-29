package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 88.5H302.627l72.5-72.5h-45.254l-72.5 72.5h-2.246l-72.5-72.5h-45.254l72.5 72.5H40a24.028 24.028 0 0 0-24 24v296a24.028 24.028 0 0 0 24 24h112V496h224v-63.5h96a24.028 24.028 0 0 0 24-24v-296a24.028 24.028 0 0 0-24-24ZM344 464H184v-31.5h160Zm120-63.5H48v-280h416Z"/>`),
		g.Group(children),
	)
}