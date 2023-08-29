package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m330 16l-42.68 42.7L453.3 224.68L496 182L330 16z"/><ellipse cx="224.68" cy="287.3" fill="none" rx="20.03" ry="19.96"/><path fill="currentColor" d="M429.21 243.85L268 82.59L249.65 168L16 402l94 94l234.23-233.8Zm-189 56.07a20 20 0 1 1 0-25.25a20 20 0 0 1-.02 25.25Z"/>`),
		g.Group(children),
	)
}