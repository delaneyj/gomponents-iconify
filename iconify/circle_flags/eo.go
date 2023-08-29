package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<clipPath id="svgIDa"><circle cx="256" cy="256" r="256"/></clipPath><g clip-path="url(#svgIDa)"><path fill="#6da544" d="M0 256L256 0h256v512H0z"/><path fill="#eee" d="M0 0h256v256H0z"/><path fill="#6da544" d="m152.4 89l16.6 51h53.6l-43.4 31.6l16.6 51l-43.4-31.5l-43.4 31.5l16.6-51L82.2 140h53.6z"/></g>`),
		g.Group(children),
	)
}