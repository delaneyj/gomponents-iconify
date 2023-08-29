package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActionRedo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M361.376 495.163L226.753 360.54l22.627-22.627l111.996 111.996l111.997-111.996L496 360.54L361.376 495.163z"/><path fill="currentColor" d="M377.377 472.52h-32V196.426C345.377 114.584 278.794 48 196.952 48c-83.229 0-148.426 63.106-148.426 143.667h-32c0-48.024 18.85-92.569 53.079-125.429C103.35 33.842 148.576 16 196.952 16c99.487 0 180.425 80.938 180.425 180.426Z"/>`),
		g.Group(children),
	)
}