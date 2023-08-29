package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageUpgradeOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.243 20.243l-2.829-2.829L18 16l-4.249 4.249l1.415 1.414L18 18.829l2.828 2.828l1.415-1.414zM7 15h4v2H7z"/><path fill="currentColor" d="M22 6H2v2h2v11a2.003 2.003 0 0 0 2 2h5l7-7l2 2V8h2ZM6 8h12v3l-8 8H6Z"/>`),
		g.Group(children),
	)
}