package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayOrPauseButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="m13 16l19 16l-19 16zm22.8 0h5.7v32h-5.7zm9.5 0H51v32h-5.7z"/>`),
		g.Group(children),
	)
}