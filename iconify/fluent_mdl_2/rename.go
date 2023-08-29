package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rename(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M255 1149V895h770v254H255zm1026 260H-1V639h1282v130H129v510h1152v130zm768-770v770h-386v-130h256V769h-256V639h386zm-512 928q111 96 255 96h1v130h-1q-176 0-318-110q-70 54-151 82t-171 28h-1v-130h1q145 0 255-93V474q-52-44-116-66t-139-23h-1V255h1q91 0 172 27t150 79q138-106 318-106h1v130h-1q-148 0-255 92v1090z"/>`),
		g.Group(children),
	)
}