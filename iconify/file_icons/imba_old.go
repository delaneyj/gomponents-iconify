package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImbaOld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M320 0H192L0 192v128l192 192h128l192-192V192L320 0zm-64.002 346.75L165.25 256L256 165.25L346.75 256l-90.752 90.75z"/>`),
		g.Group(children),
	)
}