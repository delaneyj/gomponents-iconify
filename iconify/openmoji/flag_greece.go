package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagGreece(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M5 17h62v38H5z"/><path fill="#1e50a0" d="M5 34h62v4H5zm0-8.25h62v4H5zm0 16.5h62v4H5zM5 50h62v5H5zm0-33h62v5H5z"/><path fill="#1e50a0" d="M5 17h22v21H5z"/><path fill="#fff" d="M14.5 17h4v22h-4z"/><path fill="#fff" d="M5 25.75h22v4H5z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}