package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iran(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#fff" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="#46b000" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v6.17h7.37A45.73 45.73 0 0 1 24 16.84a45.62 45.62 0 0 0 14.65 2.38h7.37V13a.94.94 0 0 0-1.02-.75Z"/><path fill="#e04122" d="M24 30.5a45.72 45.72 0 0 0-14.63-2.37H2v6.17a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79v-6.17h-7.32A45.62 45.62 0 0 1 24 30.5Z"/><path fill="none" stroke="#46b000" stroke-linecap="round" stroke-linejoin="round" d="M2 15.93h7.37A45.73 45.73 0 0 1 24 18.31a45.61 45.61 0 0 0 14.65 2.37h7.37"/><path fill="none" stroke="#e04122" stroke-linecap="round" stroke-linejoin="round" d="M2 26.76h7.37A45.73 45.73 0 0 1 24 29.14a45.89 45.89 0 0 0 14.65 2.37h7.37"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="none" stroke="#e04122" stroke-linecap="round" stroke-linejoin="round" d="M22.47 26.36a2.58 2.58 0 0 0 3.39-1.76a3.17 3.17 0 0 0-.37-2.6"/><path fill="none" stroke="#e04122" stroke-linecap="round" stroke-linejoin="round" d="M25.46 26.94a4 4 0 0 1-3.38-3.07a2.67 2.67 0 0 1 .36-2.48m1.52-.28v5.76m2.46-5.52s2.31 1.59.61 5M21.5 20.4s-2.31.68-.61 4.75M23 20.47a5.63 5.63 0 0 0 .94.27a5 5 0 0 0 .94.1"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}