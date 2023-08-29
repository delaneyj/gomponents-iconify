package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpikesInit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16c-7.5 67.5-37.5 150-37.5 180c0 15 15 30 37.5 30s37.5-15 37.5-30c0-30-30-112.5-37.5-180zm-60 202.5c-30 0-112.5 30-180 37.5c67.5 7.5 150 37.5 180 37.5c15 0 30-15 30-37.5s-15-37.5-30-37.5zm120 0c-15 0-30 15-30 37.5s15 37.5 30 37.5c30 0 112.5-30 180-37.5c-67.5-7.5-150-37.5-180-37.5zM256 286c-22.5 0-37.5 15-37.5 30c0 30 30 112.5 37.5 180c7.5-67.5 37.5-150 37.5-180c0-15-15-30-37.5-30z"/>`),
		g.Group(children),
	)
}