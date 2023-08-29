package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotSurface(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M106 53.5c45 150-45 150 0 300h60c-45-150 45-150 0-300h-60zm120 0c45 150-45 150 0 300h60c-45-150 45-150 0-300h-60zm120 0c45 150-45 150 0 300h60c-45-150 45-150 0-300h-60zm-270 315c-15 0-30 30-30 30c-30 0-30 0-30 30v30h480v-30c0-30 0-30-30-30c0 0-15-30-30-30c-30 0-30 45-60 45s-30-45-60-45s-30 45-60 45s-30-45-60-45s-30 45-60 45s-30-45-60-45z"/>`),
		g.Group(children),
	)
}