package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Russia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#e04122" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 11.75h-6.32A45.89 45.89 0 0 1 24 9.38A45.73 45.73 0 0 0 9.37 7H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79v-26a.94.94 0 0 0-1-.84Z"/><path fill="#009fd9" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M38.68 20.68A45.89 45.89 0 0 1 24 18.31a45.73 45.73 0 0 0-14.63-2.38H2v9.73h7.37A46 46 0 0 1 24 28a45.62 45.62 0 0 0 14.65 2.38h7.37v-9.7Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 11.75h-6.32A45.89 45.89 0 0 1 24 9.38A45.73 45.73 0 0 0 9.37 7H3c-.58 0-1 .35-1 .79v8.14h7.37A45.73 45.73 0 0 1 24 18.31a45.89 45.89 0 0 0 14.65 2.37h7.37v-8.14a.94.94 0 0 0-1.02-.79Z"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}