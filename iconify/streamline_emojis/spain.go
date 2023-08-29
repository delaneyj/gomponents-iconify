package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M13 26.35c.89.07 1.77.16 2.65.28a1.8 1.8 0 0 0 2-1.79v-4.39a47.67 47.67 0 0 0-6.34-.71v4.81a1.8 1.8 0 0 0 1.69 1.8Z"/><path fill="#e04122" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M46.05 33.33h-7.37A45.61 45.61 0 0 1 24 31a45.73 45.73 0 0 0-14.63-2.42H2v5.72a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79Zm0-20.33a.94.94 0 0 0-1-.79h-6.37A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79V14h7.37A46 46 0 0 1 24 16.38a45.62 45.62 0 0 0 14.65 2.38h7.37Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M8.25 19.68h1.51v5.96H8.25zm12.52 7.36l-1.5-.39v-5.96l1.5.39v5.96z"/><path fill="#e04122" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.63 26.63a1.8 1.8 0 0 0 2-1.79v-1.33c-1-.18-2.07-.33-3.11-.45v3.44ZM14.57 20c-1.07-.12-2.15-.2-3.23-.25v3.05c1.08 0 2.16.13 3.23.25Zm-1.52-1.7a23.86 23.86 0 0 1 3.18.39s1.31-.43.85-1.36c-.67-1.33-2.49-.55-2.49-.55s-1.41-1.09-2.28-.26s.74 1.78.74 1.78Zm4.63 4.48a3.44 3.44 0 0 1 3.09.4v2a3.44 3.44 0 0 0-3.09-.4Zm-9.43-1.04a3.44 3.44 0 0 1 3.09.4v2a3.44 3.44 0 0 0-3.09-.4Z"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}