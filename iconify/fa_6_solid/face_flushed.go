package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceFlushed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0zm176 128c0 8.8 7.2 16 16 16h128c8.8 0 16-7.2 16-16s-7.2-16-16-16H192c-8.8 0-16 7.2-16 16zm-16-88a72 72 0 1 0 0-144a72 72 0 1 0 0 144zm264-72a72 72 0 1 0-144 0a72 72 0 1 0 144 0zm-288 0a24 24 0 1 1 48 0a24 24 0 1 1-48 0zm192 0a24 24 0 1 1 48 0a24 24 0 1 1-48 0z"/>`),
		g.Group(children),
	)
}