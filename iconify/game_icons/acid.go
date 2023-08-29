package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256.875 16A30 30 0 0 0 226 46a30 30 0 0 0 60 0a30 30 0 0 0-29.125-30zm-45 75A30 30 0 0 0 181 121a30 30 0 0 0 60 0a30 30 0 0 0-29.125-30zm74.563 30A15 15 0 0 0 271 136a15 15 0 0 0 30 0a15 15 0 0 0-14.563-15zm-30 45A15 15 0 0 0 241 181a15 15 0 0 0 30 0a15 15 0 0 0-14.563-15zM196 196c-45 0-15 30 0 45c0 150-120 225-120 255h360c0-30-120-105-120-255c15-15 45-45 0-45c-15 0-30 15-60 15s-45-15-60-15z"/>`),
		g.Group(children),
	)
}