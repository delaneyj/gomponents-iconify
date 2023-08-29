package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M464 256a208 208 0 1 0-416 0a208 208 0 1 0 416 0zM0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0zm306.7 69.1l-144.3 55.5c-19.4 7.5-38.5-11.6-31-31l55.5-144.3c3.3-8.5 9.9-15.1 18.4-18.4l144.3-55.5c19.4-7.5 38.5 11.6 31 31l-55.5 144.3c-3.2 8.5-9.9 15.1-18.4 18.4zM288 256a32 32 0 1 0-64 0a32 32 0 1 0 64 0z"/>`),
		g.Group(children),
	)
}