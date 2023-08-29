package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayPauseBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M188 64v128a12 12 0 0 1-24 0V64a12 12 0 0 1 24 0Zm36-12a12 12 0 0 0-12 12v128a12 12 0 0 0 24 0V64a12 12 0 0 0-12-12Zm-76 76a19.71 19.71 0 0 1-9.19 16.71l-88.18 56.16A19.91 19.91 0 0 1 20 184.15V71.85a19.91 19.91 0 0 1 30.63-16.72l88.18 56.16A19.71 19.71 0 0 1 148 128Zm-27.62 0L44 79.37v97.26Z"/>`),
		g.Group(children),
	)
}