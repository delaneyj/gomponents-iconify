package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBahamas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#00ABC9" d="M1.364 29.987A3.966 3.966 0 0 0 4 31h28a4 4 0 0 0 4-4v-4.5H11.442L1.364 29.987z"/><path fill="#FAE042" d="m17.5 18l-6.058 4.5H36v-9H11.442z"/><path fill="#00ABC9" d="M32 5H4c-1.015 0-1.931.39-2.636 1.013L11.442 13.5H36V9a4 4 0 0 0-4-4z"/><path fill="#141414" d="m17.5 18l-6.058-4.5L1.364 6.013A3.974 3.974 0 0 0 0 9v18c0 1.194.534 2.254 1.364 2.987L11.442 22.5L17.5 18z"/>`),
		g.Group(children),
	)
}