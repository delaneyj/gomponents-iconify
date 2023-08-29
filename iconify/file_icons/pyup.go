package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pyup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 128.083v253.776l89.629 52.571l.004-255.146l132.247-75.759l63.99 38.222V36.696L221.704 0L0 128.083zm354.21 203.912l-132.332 73.393l-63.33-33.031v102.497L221.877 512L443 381.858V126.565L354.21 75.83v256.166zm-132.347-3.113l64.006-36.629v-74.58l-64.345-35.65l-63.64 36.233v74.19l63.98 36.436z"/>`),
		g.Group(children),
	)
}