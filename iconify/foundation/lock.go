package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M82.105 44.218h-8.858v-8.431c.003-.036.003-.071.003-.102c0-13.073-10.636-23.71-23.713-23.71c-13.073 0-23.71 10.637-23.71 23.71v8.533h-7.931a2.62 2.62 0 0 0-2.621 2.621v38.565a2.62 2.62 0 0 0 2.621 2.621h64.21a2.62 2.62 0 0 0 2.621-2.621V46.839a2.621 2.621 0 0 0-2.622-2.621zm-42.314-8.533c0-5.375 4.371-9.741 9.746-9.741c5.341 0 9.695 4.32 9.747 9.649l-.003.031h.003v8.594H39.791v-8.533z"/>`),
		g.Group(children),
	)
}