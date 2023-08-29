package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NrFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagNr4x30"><path fill-opacity=".7" d="M-54.7 0H628v512H-54.7z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagNr4x30)" transform="translate(51.3) scale(.9375)"><path fill="#002170" d="M-140 0H884v512H-140z"/><path fill="#ffb20d" d="M-140 234.1H884V278H-140z"/><path fill="#fff" d="m161.8 438l-33-33l-10.5 45.4l-12-45l-31.9 34l12.1-45L42 407.9l33-33l-45.4-10.6l45-12l-34-31.8l45 12L72 288l33 33l10.6-45.4l12 45l31.8-34l-12 45l44.5-13.5l-33 33l45.4 10.5l-45 12l34 32l-45-12.2z"/></g>`),
		g.Group(children),
	)
}