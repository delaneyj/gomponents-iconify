package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autoinstallation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.424 16.5h5.12v-6h4.448a9.004 9.004 0 1 0-6.458 9.11Zm-1.908-5.526a2.5 2.5 0 1 1 2.5 2.5a2.5 2.5 0 0 1-2.5-2.5Z"/><path fill="currentColor" d="M22.045 18h-3v-6h-2v6h-3l4 4l4-4z"/>`),
		g.Group(children),
	)
}