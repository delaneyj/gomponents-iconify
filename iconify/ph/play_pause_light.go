package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayPauseLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M182 64v128a6 6 0 0 1-12 0V64a6 6 0 0 1 12 0Zm42-6a6 6 0 0 0-6 6v128a6 6 0 0 0 12 0V64a6 6 0 0 0-6-6Zm-82 70a13.77 13.77 0 0 1-6.41 11.65L47.41 195.8A13.91 13.91 0 0 1 26 184.15V71.85A13.91 13.91 0 0 1 47.41 60.2l88.18 56.15A13.77 13.77 0 0 1 142 128Zm-12 0a1.77 1.77 0 0 0-.85-1.53L41 70.32a1.87 1.87 0 0 0-1-.32a2.13 2.13 0 0 0-1 .25a1.76 1.76 0 0 0-1 1.6v112.3a1.76 1.76 0 0 0 1 1.6a1.9 1.9 0 0 0 2-.07l88.19-56.15A1.77 1.77 0 0 0 130 128Z"/>`),
		g.Group(children),
	)
}