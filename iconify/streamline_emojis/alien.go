package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alien(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#00dba8" d="M44 19c0 11.05-13.5 22.5-20 22.5S4 30.05 4 19S13 1.5 24 1.5S44 8 44 19Z"/><path fill="#00ad85" d="M24 38.39c-6.22 0-18.82-10.47-19.91-21C4 17.89 4 18.43 4 19c0 11.05 13.5 22.5 20 22.5S44 30.05 44 19c0-.57 0-1.11-.08-1.65c-1.1 10.57-13.7 21.04-19.92 21.04Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" d="M29 30a9 9 0 0 1-10 0"/><path fill="#45413c" d="M8 45.5a16 1.5 0 1 0 32 0a16 1.5 0 1 0-32 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M44 19c0 11.05-13.5 22.5-20 22.5S4 30.05 4 19S13 1.5 24 1.5S44 8 44 19Z"/><path fill="#bf8df2" d="M10.88 12.41c-2.37 1.81-2.44 5.7-.17 8.69s6.28 3.45 9 3.06a8.78 8.78 0 0 0-.41-9.61c-2.3-2.99-6.04-3.94-8.42-2.14Z"/><path fill="#dabff5" d="M19.3 14.55c-2.27-3-6-3.94-8.42-2.14a5.3 5.3 0 0 0-1.67 5.44a4.64 4.64 0 0 1 1.67-2.71c2.38-1.81 6.15-.85 8.42 2.13A8.24 8.24 0 0 1 20.85 21a8.71 8.71 0 0 0-1.55-6.45Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M10.88 12.41c-2.37 1.81-2.44 5.7-.17 8.69s6.28 3.45 9 3.06a8.78 8.78 0 0 0-.41-9.61c-2.3-2.99-6.04-3.94-8.42-2.14Z"/><path fill="#bf8df2" d="M37.12 12.41c2.37 1.81 2.44 5.7.17 8.69s-6.28 3.45-9 3.06a8.78 8.78 0 0 1 .41-9.61c2.3-2.99 6.04-3.94 8.42-2.14Z"/><path fill="#dabff5" d="M28.7 14.55c2.27-3 6-3.94 8.42-2.14a5.3 5.3 0 0 1 1.67 5.44a4.64 4.64 0 0 0-1.67-2.71c-2.38-1.81-6.15-.85-8.42 2.13A8.24 8.24 0 0 0 27.15 21a8.71 8.71 0 0 1 1.55-6.45Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M37.12 12.41c2.37 1.81 2.44 5.7.17 8.69s-6.28 3.45-9 3.06a8.78 8.78 0 0 1 .41-9.61c2.3-2.99 6.04-3.94 8.42-2.14ZM25.5 26v1.5m-3-1.5v1.5"/><path fill="#00ad85" d="M10.81 28.17a2.5 1.5 0 1 0 5 0a2.5 1.5 0 1 0-5 0Zm21.38 0a2.5 1.5 0 1 0 5 0a2.5 1.5 0 1 0-5 0Z"/><path fill="#00f5bc" d="M18.04 5.39a6 1.5 0 1 0 12 0a6 1.5 0 1 0-12 0Z"/>`),
		g.Group(children),
	)
}