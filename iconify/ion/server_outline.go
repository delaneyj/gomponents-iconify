package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<ellipse cx="256" cy="128" fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="32" rx="192" ry="80"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="32" d="M448 214c0 44.18-86 80-192 80S64 258.18 64 214m384 86c0 44.18-86 80-192 80S64 344.18 64 300"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="32" d="M64 127.24v257.52C64 428.52 150 464 256 464s192-35.48 192-79.24V127.24"/>`),
		g.Group(children),
	)
}