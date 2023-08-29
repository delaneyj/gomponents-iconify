package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22H5V2h14Zm0-22H5a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2Z"/><path fill="currentColor" d="M8 4a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm10 0a1 1 0 1 1-1-1a1 1 0 0 1 1 1ZM8 20a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm10 0a1 1 0 1 1-1-1a1 1 0 0 1 1 1ZM12 6a5 5 0 0 0-5 5a4.936 4.936 0 0 0 1 3h3v1.9a5.001 5.001 0 0 0 6-4.9c0-5-5-5-5-5Zm0 7a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}