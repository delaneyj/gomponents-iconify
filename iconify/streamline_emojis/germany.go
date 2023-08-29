package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Germany(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M2 26.16h7.37A46 46 0 0 1 24 28.53h0a45.62 45.62 0 0 0 14.65 2.38h7.37v-9.73h-7.34A45.89 45.89 0 0 1 24 18.81h0a45.73 45.73 0 0 0-14.63-2.38H2Z"/><path fill="#656769" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v8.14h7.37A45.73 45.73 0 0 1 24 18.81a45.89 45.89 0 0 0 14.65 2.37h7.37V13a.94.94 0 0 0-1.02-.75Z"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}