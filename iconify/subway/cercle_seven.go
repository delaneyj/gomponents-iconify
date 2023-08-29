package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CercleSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0zm0 64c35.4 0 64 28.6 64 64s-28.6 64-64 64s-64-28.6-64-64s28.6-64 64-64zm85.3 384H170.7v-21.3H192V256h-21.3v-21.3H320v192h21.3V448z"/>`),
		g.Group(children),
	)
}