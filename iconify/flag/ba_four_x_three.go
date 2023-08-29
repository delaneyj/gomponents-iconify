package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagBa4x30"><path fill-opacity=".7" d="M-85.3 0h682.6v512H-85.3z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagBa4x30)" transform="translate(80) scale(.9375)"><path fill="#009" d="M-85.3 0h682.6v512H-85.3V0z"/><path fill="#FC0" d="m56.5 0l511 512.3V.3L56.5 0z"/><path fill="#FFF" d="M439.9 481.5L412 461.2l-28.6 20.2l10.8-33.2l-28.2-20.5h35l10.8-33.2l10.7 33.3h35l-28 20.7l10.4 33zm81.3 10.4l-35-.1l-10.7-33.3l-10.8 33.2h-35l28.2 20.5l-10.8 33.2l28.6-20.2l28 20.3l-10.5-33l28-20.6zM365.6 384.7l28-20.7l-35-.1l-10.7-33.2l-10.8 33.2l-35-.1l28.2 20.5l-10.8 33.3l28.6-20.3l28 20.4l-10.5-33zm-64.3-64.5l28-20.6l-35-.1l-10.7-33.3l-10.9 33.2h-34.9l28.2 20.5l-10.8 33.2l28.6-20.2l27.9 20.3l-10.4-33zm-63.7-63.6l28-20.7h-35L220 202.5l-10.8 33.2h-35l28.2 20.4l-10.8 33.3l28.6-20.3l28 20.4l-10.5-33zm-64.4-64.3l28-20.6l-35-.1l-10.7-33.3l-10.9 33.2h-34.9L138 192l-10.8 33.2l28.6-20.2l27.9 20.3l-10.4-33zm-63.6-63.9l27.9-20.7h-35L91.9 74.3L81 107.6H46L74.4 128l-10.9 33.2L92.1 141l27.8 20.4l-10.3-33zm-64-64l27.9-20.7h-35L27.9 10.3L17 43.6h-35L10.4 64l-11 33.3L28.1 77l27.8 20.4l-10.3-33zm-64-64L9.4-20.3h-35l-10.7-33.3L-47-20.4h-35L-53.7 0l-10.8 33.2L-35.9 13l27.8 20.4l-10.3-33z"/></g>`),
		g.Group(children),
	)
}