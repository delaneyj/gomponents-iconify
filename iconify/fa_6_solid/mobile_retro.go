package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileRetro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 64C0 28.7 28.7 0 64 0h192c35.3 0 64 28.7 64 64v384c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V64zm64 96v64c0 17.7 14.3 32 32 32h128c17.7 0 32-14.3 32-32v-64c0-17.7-14.3-32-32-32H96c-17.7 0-32 14.3-32 32zm16 192a24 24 0 1 0 0-48a24 24 0 1 0 0 48zm24 56a24 24 0 1 0-48 0a24 24 0 1 0 48 0zm56-56a24 24 0 1 0 0-48a24 24 0 1 0 0 48zm24 56a24 24 0 1 0-48 0a24 24 0 1 0 48 0zm56-56a24 24 0 1 0 0-48a24 24 0 1 0 0 48zm24 56a24 24 0 1 0-48 0a24 24 0 1 0 48 0zM128 48c-8.8 0-16 7.2-16 16s7.2 16 16 16h64c8.8 0 16-7.2 16-16s-7.2-16-16-16h-64z"/>`),
		g.Group(children),
	)
}