package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Postbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M11.71 44.4a12.29 1.6 0 1 0 24.58 0a12.29 1.6 0 1 0-24.58 0Z" opacity=".15"/><path fill="#ff6242" d="M10.81 2.9H37.2v31H10.81Z"/><path fill="#ff866e" d="M34 2.9H14a3.2 3.2 0 0 0-3.2 3.2v2.48A3.21 3.21 0 0 1 14 5.37h20a3.21 3.21 0 0 1 3.2 3.21V6.1A3.2 3.2 0 0 0 34 2.9Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10.81 2.9H37.2v31H10.81Z"/><path fill="#debb7e" d="M21.76 33.89h4.47v9.54a1.07 1.07 0 0 1-1.07 1.07h-2.33a1.07 1.07 0 0 1-1.07-1.07v-9.54Z"/><path fill="#b89558" d="M21.76 33.89h4.47v2.72h-4.47z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.76 33.89h4.47v9.54a1.07 1.07 0 0 1-1.07 1.07h-2.33a1.07 1.07 0 0 1-1.07-1.07v-9.54h0Z"/><path fill="#656769" d="M14.41 11.06h19.17v3.47H14.41Z"/><path fill="#525252" d="M32.52 11.06h-17a1.07 1.07 0 0 0-1.07 1.07v1.3a1.07 1.07 0 0 1 1.07-1.07h17a1.07 1.07 0 0 1 1.07 1.07v-1.3a1.07 1.07 0 0 0-1.07-1.07Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.41 11.06h19.17v3.47H14.41Z"/><path fill="#e04122" d="M33.9 6.58H14.1a1.12 1.12 0 0 0 0 2.23h19.8a1.12 1.12 0 0 0 0-2.23Z"/><path fill="#ff6242" d="M14.1 7.7h19.8a1.11 1.11 0 0 1 1 .55a1 1 0 0 0 .1-.55a1.12 1.12 0 0 0-1.1-1.12H14.1A1.12 1.12 0 0 0 13 7.7a1 1 0 0 0 .15.55a1.11 1.11 0 0 1 .95-.55Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.9 6.58H14.1a1.12 1.12 0 0 0 0 2.23h19.8a1.12 1.12 0 0 0 0-2.23Z"/><path fill="#fff" d="M33.56 21.16h-9.79a1.07 1.07 0 0 1-1-.63l-4.26-9.47h11.33l4.21 9.34a.54.54 0 0 1-.49.76Z"/><path fill="#f0f0f0" d="m30.87 13.35l-1.03-2.29H18.53l1.03 2.29h11.31z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.56 21.16h-9.79a1.07 1.07 0 0 1-1-.63l-4.26-9.47h11.33l4.21 9.34a.54.54 0 0 1-.49.76Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m34.03 20.87l-7.57-6.13l.85-3.68m.3 4.61l-4.36 5.35m3.46-7.35l-5.24-2.61"/>`),
		g.Group(children),
	)
}