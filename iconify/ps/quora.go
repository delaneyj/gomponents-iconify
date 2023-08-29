package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M316 390q44-27 70.5-74T413 214q0-86-60-147T208 6T63 67T3 214t60 146.5T208 421q26 0 48-6q35 64 112 49v-35q-40-11-52-39zm-2-149q0 52-29 89q-40-44-102-41v41q30 1 51 37q-12 3-24 3q-44 0-74.5-37.5T105 241v-54q0-53 31-91t74-38t73.5 38t30.5 91v54z"/>`),
		g.Group(children),
	)
}