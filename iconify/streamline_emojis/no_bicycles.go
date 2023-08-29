package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoBicycles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#656769" d="M8.4 24a15.6 15.6 0 1 0 31.2 0a15.6 15.6 0 1 0-31.2 0Z"/><path fill="#525252" d="M24 11.27a15.61 15.61 0 0 1 15.53 14.17c0-.48.07-1 .07-1.44a15.6 15.6 0 0 0-31.2 0c0 .48 0 1 .07 1.44A15.61 15.61 0 0 1 24 11.27Z"/><path fill="none" stroke="#fff" d="M11.51 27.65a5.01 5.01 0 1 0 10.02 0a5.01 5.01 0 1 0-10.02 0Zm15.14 0a5.01 5.01 0 1 0 10.02 0a5.01 5.01 0 1 0-10.02 0Z"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m16.52 27.65l2.53-7.6v-4.71h-4.12m4.12 4.71h10.76l-5.38 7.6l-5.38-7.6z"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M24.43 27.65h7.23l-1.85-7.6v-2.12m-2.35 0h4.71"/><path fill="#fff" d="M13.99 27.65a2.53 2.53 0 1 0 5.06 0a2.53 2.53 0 1 0-5.06 0Zm15.14 0a2.53 2.53 0 1 0 5.06 0a2.53 2.53 0 1 0-5.06 0Z"/><path fill="#45413c" d="M11.98 45.43a12.02 1.57 0 1 0 24.04 0a12.02 1.57 0 1 0-24.04 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M8.4 24a15.6 15.6 0 1 0 31.2 0a15.6 15.6 0 1 0-31.2 0Z"/><path fill="#ff6242" d="M24 4.15A19.85 19.85 0 1 0 43.85 24A19.85 19.85 0 0 0 24 4.15Zm0 4.25a15.55 15.55 0 0 1 15.17 19.18L11 15.41A15.57 15.57 0 0 1 24 8.4Zm0 31.2A15.55 15.55 0 0 1 8.83 20.42L37 32.59a15.57 15.57 0 0 1-13 7.01Z"/><path fill="#ff866e" d="M24 6.65a19.84 19.84 0 0 1 19.79 18.6c0-.41.06-.83.06-1.25a19.85 19.85 0 0 0-39.7 0c0 .42 0 .84.06 1.25A19.84 19.84 0 0 1 24 6.65Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 4.15A19.85 19.85 0 1 0 43.85 24A19.85 19.85 0 0 0 24 4.15Zm0 4.25a15.55 15.55 0 0 1 15.17 19.18L11 15.41A15.57 15.57 0 0 1 24 8.4Zm0 31.2A15.55 15.55 0 0 1 8.83 20.42L37 32.59a15.57 15.57 0 0 1-13 7.01Z"/>`),
		g.Group(children),
	)
}