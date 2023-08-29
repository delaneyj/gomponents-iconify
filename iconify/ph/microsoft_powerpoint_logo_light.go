package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftPowerpointLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M96 98H80a6 6 0 0 0-6 6v48a6 6 0 0 0 12 0v-10h10a22 22 0 0 0 0-44Zm0 32H86v-20h10a10 10 0 0 1 0 20Zm40-104a102.35 102.35 0 0 0-81 40H40a14 14 0 0 0-14 14v96a14 14 0 0 0 14 14h15a102 102 0 1 0 81-164Zm89.8 96H158V80a14 14 0 0 0-14-14h-2V38.2a90.15 90.15 0 0 1 83.8 83.8ZM130 38.21V66H70.78A90.39 90.39 0 0 1 130 38.21ZM38 176V80a2 2 0 0 1 2-2h104a2 2 0 0 1 2 2v96a2 2 0 0 1-2 2H40a2 2 0 0 1-2-2Zm32.78 14H130v27.79A90.39 90.39 0 0 1 70.78 190ZM142 217.8V190h2a14 14 0 0 0 14-14v-42h67.8a90.14 90.14 0 0 1-83.8 83.8Z"/>`),
		g.Group(children),
	)
}