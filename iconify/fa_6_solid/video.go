package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 128c0-35.3 28.7-64 64-64h256c35.3 0 64 28.7 64 64v256c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V128zm559.1-28.2c10.4 5.6 16.9 16.4 16.9 28.2v256c0 11.8-6.5 22.6-16.9 28.2s-23 5-32.9-1.6l-96-64l-14.2-9.5V174.9l14.2-9.5l96-64c9.8-6.5 22.4-7.2 32.9-1.6z"/>`),
		g.Group(children),
	)
}