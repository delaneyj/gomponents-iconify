package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeLightZeit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs><linearGradient id="vscodeIconsFileTypeLightZeit0" x1="1.288" x2="1.143" y1="32.55" y2="32.75" gradientTransform="matrix(114 0 0 -100 -113 3301)" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#fff"/><stop offset="1"/></linearGradient></defs><path fill="url(#vscodeIconsFileTypeLightZeit0)" fill-rule="evenodd" d="m16 3.719l14 24.562H2L16 3.719z"/>`),
		g.Group(children),
	)
}