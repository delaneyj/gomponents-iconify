package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AgFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagAg4x30"><path fill-opacity=".7" d="M-79.7 0H603v512H-79.7z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagAg4x30)" transform="translate(74.7) scale(.9375)"><path fill="#fff" d="M-79.7 0H603v512H-79.7V0Z"/><path d="M-79.6 0H603v204.8H-79.7L-79.6 0Z"/><path fill="#0072c6" d="M21.3 203.2h480v112h-480v-112Z"/><path fill="#ce1126" d="M603 .1V512H261.6L603 0v.1ZM-79.7.1V512h341.3L-79.7 0v.1Z"/><path fill="#fcd116" d="M440.4 203.3L364 184l64.9-49l-79.7 11.4l41-69.5l-70.7 41L332.3 37l-47.9 63.8l-19.3-74l-21.7 76.3l-47.8-65l13.7 83.2L138.5 78l41 69.5l-77.4-12.5l63.8 47.8L86 203.3h354.3z"/></g>`),
		g.Group(children),
	)
}