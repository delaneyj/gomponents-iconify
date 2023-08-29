package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pakistan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#46b000" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M13.38 7.66c-1.29-.1-2.59-.16-3.9-.16H3.15c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.33c1.31 0 2.61.06 3.9.16Z"/><path fill="#fff" d="m35.15 18.42l.93-.46a.25.25 0 0 0 .14-.19l.15-1a.25.25 0 0 1 .43-.14l.73.74a.27.27 0 0 0 .22.08l1-.18a.25.25 0 0 1 .27.37l-.49.92a.25.25 0 0 0 0 .24l.49.92a.25.25 0 0 1-.27.36l-1-.17a.23.23 0 0 0-.22.07l-.73.75a.25.25 0 0 1-.43-.15l-.15-1a.22.22 0 0 0-.14-.19l-.93-.47a.25.25 0 0 1 0-.5Zm.75 12.63c-3.73 1.89-8.76.08-11.24-4s-1.47-9 2.26-10.88a6.87 6.87 0 0 1 2.08-.69a8.54 8.54 0 0 0-5.3.73c-4.32 2.19-5.49 7.85-2.62 12.63s8.72 6.88 13 4.68a7.44 7.44 0 0 0 3.5-3.72a6.29 6.29 0 0 1-1.68 1.25Z"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}