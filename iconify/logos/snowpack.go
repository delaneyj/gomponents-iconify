package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snowpack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M253.964 185.078L138.764 5.88A12.802 12.802 0 0 0 127.998 0a12.802 12.802 0 0 0-10.768 5.88L2.03 185.078a12.8 12.8 0 0 0-.468 13.056a12.803 12.803 0 0 0 11.236 6.664h230.398c4.684 0 8.992-2.556 11.236-6.668a12.793 12.793 0 0 0-.468-13.052ZM127.997 36.472l34.156 53.127h-34.156l-25.6 25.6l-15.224-15.224l40.824-63.503Z"/>`),
		g.Group(children),
	)
}