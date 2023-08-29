package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="url(#f691id0)" d="M16 30.137c7.732 0 14-3.643 14-8.137c0-4.494-6.268-8.137-14-8.137S2 17.506 2 22c0 4.494 6.268 8.137 14 8.137Z"/><path fill="url(#f691id1)" d="M16 30.067c7.185 0 13.01-3.157 13.01-7.051s-5.825-7.05-13.01-7.05s-13.01 3.156-13.01 7.05c0 3.894 5.825 7.05 13.01 7.05Z"/><defs><radialGradient id="f691id0" cx="0" cy="0" r="1" gradientTransform="matrix(4.75001 -11.87503 12.18115 4.87246 23.25 15.613)" gradientUnits="userSpaceOnUse"><stop stop-color="#D8D1DF"/><stop offset=".989" stop-color="#90829F"/></radialGradient><radialGradient id="f691id1" cx="0" cy="0" r="1" gradientTransform="matrix(9.74999 -19.62496 54.14327 26.89922 10.5 31.738)" gradientUnits="userSpaceOnUse"><stop offset=".006" stop-color="#37275E"/><stop offset="1"/></radialGradient></defs></g>`),
		g.Group(children),
	)
}