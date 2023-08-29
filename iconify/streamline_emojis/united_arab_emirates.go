package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnitedArabEmirates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="#46b000" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v8.29h7.37A45.72 45.72 0 0 1 24 19a45.62 45.62 0 0 0 14.65 2.38h7.37V13a.94.94 0 0 0-1.02-.75Z"/><path fill="#656769" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M38.68 30.54A45.89 45.89 0 0 1 24 28.17a45.73 45.73 0 0 0-14.63-2.38H2v8.51a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79v-8.51Z"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M13.53 7.69c-1.38-.12-2.77-.19-4.16-.19H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37c1.39 0 2.78.07 4.16.19Z"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}