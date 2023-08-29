package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TridentEmblem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#eda454" d="M57 24.6C54 17.4 44.8 8.7 44.8 8.7s3 7.4.8 17.3l3.5-3s3.3 1.4 3.7 5c.5 4.4-2 10.6-2 10.6s-5.5-1.7-13.9-1.6l-3-19.6l7.2 4.7L32 2l-9 19.9l7.2-4.7l-3 19.6c-8.4-.1-13.9 1.6-13.9 1.6s-2.5-6.2-2-10.6c.4-3.5 3.7-5 3.7-5l3.5 3c-2.2-9.8.8-17.3.8-17.3S10 17.4 7 24.6C3.9 32 2 40.8 2 40.8l6.4 5.9s5.6-1.7 10-2c4.1-.3 8.9-.4 8.9-.4l-.6 8.7l5.3 9l5.3-9l-.6-8.7s4.8.1 8.9.4c4.3.3 10 2 10 2l6.4-5.9s-1.9-8.8-5-16.2"/>`),
		g.Group(children),
	)
}