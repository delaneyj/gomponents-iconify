package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m16.08 189.4l28.58 233.87A28 28 0 0 0 72.52 448h367a28 28 0 0 0 27.86-24.73l28.54-233.87A12 12 0 0 0 484 176H28a12 12 0 0 0-11.92 13.4ZM464 124a28 28 0 0 0-28-28H244.84l-48-32H76a28 28 0 0 0-28 28v52h416Z"/>`),
		g.Group(children),
	)
}