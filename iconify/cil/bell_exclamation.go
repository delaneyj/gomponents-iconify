package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M451.671 348.569L408 267.945V184c0-83.813-68.187-152-152-152s-152 68.187-152 152v83.945l-43.671 80.623A24 24 0 0 0 81.432 384h86.944a87.762 87.762 0 0 0-.376 8a88 88 0 0 0 176 0c0-2.7-.135-5.364-.376-8h86.944a24 24 0 0 0 21.1-35.431ZM312 392a56 56 0 1 1-111.418-8h110.836a55.85 55.85 0 0 1 .582 8ZM94.863 352L136 276.055V184a120 120 0 0 1 240 0v92.055L417.137 352Z"/><path fill="currentColor" d="M240 112h32v136h-32zm0 168h32v32h-32z"/>`),
		g.Group(children),
	)
}