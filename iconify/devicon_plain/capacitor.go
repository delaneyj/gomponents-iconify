package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capacitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M108.215 0L77.512 30.691L46.96.156L27.184 19.957l80.82 80.926l19.777-19.805l-30.496-30.586L128 19.738ZM19.93 27.059L.156 46.859l30.496 30.59L0 108.191l19.715 19.813L50.43 97.246l30.547 30.535l19.777-19.8Zm0 0"/>`),
		g.Group(children),
	)
}