package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeHigh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M145.373 120H16v272h129.373l104 104H288V16h-38.627ZM128 360H48V152h80Zm128 97.373l-96-96V150.627l96-96ZM408 256a88.1 88.1 0 0 0-88-88v32a56 56 0 0 1 0 112v32a88.1 88.1 0 0 0 88-88Z"/><path fill="currentColor" d="M320 57.445v32.277C401.307 101.4 464 171.512 464 256s-62.693 154.6-144 166.278v32.277C419.005 442.66 496 358.158 496 256S419.005 69.34 320 57.445Z"/>`),
		g.Group(children),
	)
}