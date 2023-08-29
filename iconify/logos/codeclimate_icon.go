package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeclimateIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#2B2B2B" d="m84.841 28.3l28.773 28.772l28.772 29.51l27.297 27.296l-28.772 28.035l-27.297-26.559L84.84 86.582l-8.115 6.64l-36.887 37.625l-11.067 11.066L0 113.878l10.329-12.541L84.84 28.299ZM171.16 0L256 84.104l-28.772 29.51l-56.07-56.07l-19.919 19.92l-28.772-28.772L171.159 0Z"/>`),
		g.Group(children),
	)
}