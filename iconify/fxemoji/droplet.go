package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droplet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#59CAFC" d="M209.1 34.8L102 304.3c-25.8 85 22.2 174.9 107.3 200.7c85 25.8 174.9-22.2 200.7-107.3c9.6-31.7 9-63.9 0-93.4L302.8 34.8C295 8.9 267.6-5.7 241.7 2.1c-16.1 4.9-28 17.7-32.6 32.7z"/><path fill="#00B1FF" d="m102 304.3l15.7-39.5C93 353.1 143.1 445.4 231.3 472.1c57.9 17.6 117.8 2.4 160.2-34.6c-37.8 59.3-111.6 89-182.3 67.5c-85-25.8-133-115.6-107.2-200.7z"/>`),
		g.Group(children),
	)
}