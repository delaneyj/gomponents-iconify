package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Switzerland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M32.28 22.36a46.36 46.36 0 0 1-5-1v-5.64a.5.5 0 0 0-.37-.48c-1-.25-4.59-1.4-5.46-1.64a.51.51 0 0 0-.64.49v5.39a46.73 46.73 0 0 0-4.88-.95a.51.51 0 0 0-.57.5v5.46a.49.49 0 0 0 .43.49a46.36 46.36 0 0 1 5 1v4.92a.5.5 0 0 0 .37.48c1 .25 4.59 1.4 5.47 1.64a.5.5 0 0 0 .63-.49v-4.67a46.73 46.73 0 0 0 4.88.95a.51.51 0 0 0 .57-.5v-5.46a.49.49 0 0 0-.43-.49Z"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}