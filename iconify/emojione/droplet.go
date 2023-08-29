package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droplet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#75d6ff" d="M32 2C20.6 17.6 14 32 14 43.8c0 10 8.1 18.2 18 18.2s18-8.1 18-18.2C50 32 43.2 17.4 32 2z"/>`),
		g.Group(children),
	)
}