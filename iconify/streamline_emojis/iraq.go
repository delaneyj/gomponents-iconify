package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iraq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="#e04122" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M2 16.52h7.37A46 46 0 0 1 24 18.89h0a45.62 45.62 0 0 0 14.65 2.38h7.37V13a.94.94 0 0 0-1-.79h-6.34A45.89 45.89 0 0 1 24 9.88h0A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79Z"/><path fill="#525252" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M2 34.3a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46h0a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79v-8.41h-7.32A45.89 45.89 0 0 1 24 28.27h0a45.73 45.73 0 0 0-14.63-2.38H2Z"/><path fill="none" stroke="#46b000" stroke-linecap="round" stroke-linejoin="round" d="M13.75 23.81a23.08 23.08 0 0 1 6.51 1.27V23.3a1.07 1.07 0 0 0-.67-1.16c-.36-.13-.79 0-.79.49v.56a14.67 14.67 0 0 0-2.71-.54l3.23-2.86m2.57.6v4.81a4.73 4.73 0 0 1 1.38.47m-8.47-1.78V22.7m14.12-.39v5.18a15.23 15.23 0 0 1-3.37-1.07c0-1 .75-1.06 1.06-.75v-1m1.12-1.08v3.59m2.74-4.58v5.22a6.31 6.31 0 0 0 1.63.34m-6.11-5.85v-.8"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}