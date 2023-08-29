package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SprayCan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128 0h64c17.7 0 32 14.3 32 32v96H96V32c0-17.7 14.3-32 32-32zM0 256c0-53 43-96 96-96h128c53 0 96 43 96 96v208c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V256zm240 80a80 80 0 1 0-160 0a80 80 0 1 0 160 0zm16-272a32 32 0 1 1 64 0a32 32 0 1 1-64 0zm128-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64zm64 32a32 32 0 1 1 64 0a32 32 0 1 1-64 0zm32 64a32 32 0 1 1 0 64a32 32 0 1 1 0-64zm-32 128a32 32 0 1 1 64 0a32 32 0 1 1-64 0zm-64-128a32 32 0 1 1 0 64a32 32 0 1 1 0-64z"/>`),
		g.Group(children),
	)
}