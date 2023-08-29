package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMediumSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M157.65 176.1H64v159.8h93.65L288 440V72L157.65 176.1z"/><path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" stroke-width="32" d="M352 320c9.74-19.41 16-40.81 16-64c0-23.51-6-44.4-16-64m48 176c19.48-34 32-64 32-112s-12-77.7-32-112"/>`),
		g.Group(children),
	)
}