package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoadB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M288 160c0 17.7-14.3 32-32 32s-32-14.3-32-32V64c0-17.7 14.3-32 32-32s32 14.3 32 32v96z" fill="currentColor"/><path d="M256 320c-17.7 0-32 14.3-32 32v96c0 17.7 14.3 32 32 32s32-14.3 32-32v-96c0-17.7-14.3-32-32-32z" fill="currentColor"/><path d="M448 224h-96c-17.7 0-32 14.3-32 32s14.3 32 32 32h96c17.7 0 32-14.3 32-32s-14.3-32-32-32z" fill="currentColor"/><path d="M160 224H64c-17.7 0-32 14.3-32 32s14.3 32 32 32h96c17.7 0 32-14.3 32-32s-14.3-32-32-32z" fill="currentColor"/><path d="M346.5 210.7c-12.5 12.5-32.8 12.5-45.3 0s-12.5-32.8 0-45.3l67.9-67.9c12.5-12.5 32.8-12.5 45.3 0s12.5 32.8 0 45.3l-67.9 67.9z" fill="currentColor"/><path d="M210.7 301.3c-12.5-12.5-32.8-12.5-45.3 0l-67.9 67.9C85 381.7 85 402 97.5 414.5s32.8 12.5 45.3 0l67.9-67.9c12.5-12.6 12.5-32.8 0-45.3z" fill="currentColor"/><path d="M414.4 369.1l-67.9-67.9c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l67.9 67.9c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3z" fill="currentColor"/><path d="M210.7 165.5l-67.9-67.9c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l67.9 67.9c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3z" fill="currentColor"/>`),
		g.Group(children),
	)
}