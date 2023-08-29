package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChalkboardUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M160 64c0-35.3 28.7-64 64-64h352c35.3 0 64 28.7 64 64v288c0 35.3-28.7 64-64 64H336.8c-11.8-25.5-29.9-47.5-52.4-64H384v-32c0-17.7 14.3-32 32-32h64c17.7 0 32 14.3 32 32v32h64V64H224v49.1C205.2 102.2 183.3 96 160 96V64zm0 64a96 96 0 1 1 0 192a96 96 0 1 1 0-192zm-26.7 224h53.3c73.7 0 133.4 59.7 133.4 133.3c0 14.7-11.9 26.7-26.7 26.7H26.7C11.9 512 0 500.1 0 485.3C0 411.7 59.7 352 133.3 352z"/>`),
		g.Group(children),
	)
}