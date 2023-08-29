package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypePreact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#673ab8" d="m16 2l12.12 7v14L16 30L3.88 23V9z"/><ellipse cx="16" cy="16" fill="none" stroke="#fff" rx="10.72" ry="4.1" transform="rotate(-37.5 16.007 15.996)"/><ellipse cx="16" cy="16" fill="none" stroke="#fff" rx="4.1" ry="10.72" transform="rotate(-52.5 15.998 15.994)"/><circle cx="16" cy="16" r="1.86" fill="#fff"/>`),
		g.Group(children),
	)
}