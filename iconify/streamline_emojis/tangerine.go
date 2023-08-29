package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tangerine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M11.98 45.43a12.02 1.57 0 1 0 24.04 0a12.02 1.57 0 1 0-24.04 0Z" opacity=".15"/><path fill="#ffaa54" d="M7.27 28.27a16.73 16.73 0 1 0 33.46 0a16.73 16.73 0 1 0-33.46 0Z"/><path fill="#fc9" d="M24 15.72a16.74 16.74 0 0 1 16.59 14.64a18 18 0 0 0 .14-2.09a16.73 16.73 0 0 0-33.46 0a18 18 0 0 0 .14 2.09A16.74 16.74 0 0 1 24 15.72Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M7.27 28.27a16.73 16.73 0 1 0 33.46 0a16.73 16.73 0 1 0-33.46 0Z"/><path fill="#6dd627" d="M15.57 14.15s-2.66-3.27-.9-6.85s1.66-5.55 1.66-5.55s4.65 1.3 4.44 6s-5.2 6.4-5.2 6.4Z"/><path fill="#9ceb60" d="M16.27 8.93A16.31 16.31 0 0 0 18.64 3a7.24 7.24 0 0 0-2.31-1.25s.1 2-1.66 5.55s.9 6.85.9 6.85l.14-.05a6 6 0 0 1 .56-5.17Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.57 14.15s-2.66-3.27-.9-6.85s1.66-5.55 1.66-5.55s4.65 1.3 4.44 6s-5.2 6.4-5.2 6.4ZM14 17.6s.48.26 2.09-.53a2.43 2.43 0 0 0 1.6-2.07"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.57 14.15S15 10 16.42 8"/><path fill="#6dd627" d="M17 16.47a3.07 3.07 0 0 1-.88.6a8.8 8.8 0 0 1-1 .41a7 7 0 0 0-3.26-2.9c-2.52-.94-3.13 2.21-3.53-1.11s1.96-3.91 5.15-1.47A11.63 11.63 0 0 1 17 16.47Z"/><path fill="#9ceb60" d="M17 16.47a3.65 3.65 0 0 1-.78.55a11.3 11.3 0 0 0-2.7-3c-2.63-2-4.66-2-5.09 0c0-.16 0-.33-.07-.51C8 10.14 10.29 9.56 13.48 12A11.63 11.63 0 0 1 17 16.47Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M17 16.47a3.07 3.07 0 0 1-.88.6a8.8 8.8 0 0 1-1 .41a7 7 0 0 0-3.26-2.9c-2.52-.94-3.13 2.21-3.53-1.11s1.96-3.91 5.15-1.47A11.63 11.63 0 0 1 17 16.47Z"/>`),
		g.Group(children),
	)
}