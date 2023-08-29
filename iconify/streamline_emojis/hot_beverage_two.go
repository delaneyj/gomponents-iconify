package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotBeverageTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M7.15 45.42a16.85 1.58 0 1 0 33.7 0a16.85 1.58 0 1 0-33.7 0Z" opacity=".15"/><path fill="#e5f8ff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M7.25 39.08c0 4.11 8.38 6.42 16.75 6.42s16.75-2.31 16.75-6.42Z"/><path fill="#e8f4fa" d="M7.25 39.08a16.75 4.41 0 1 0 33.5 0a16.75 4.41 0 1 0-33.5 0Z"/><path fill="#fff" d="M8.33 40.62C10.71 39 16.82 37.76 24 37.76S37.29 39 39.67 40.62a2 2 0 0 0 1.08-1.54c0-2.43-7.5-4.41-16.75-4.41s-16.75 2-16.75 4.41a2 2 0 0 0 1.08 1.54Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M7.25 39.08a16.75 4.41 0 1 0 33.5 0a16.75 4.41 0 1 0-33.5 0Z"/><path fill="#e8f4fa" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.3 39.52a9.7 2.2 0 1 0 19.4 0a9.7 2.2 0 1 0-19.4 0Z"/><path fill="#e8f4fa" d="M39.43 32.91a6.17 6.17 0 1 0 0-12.34h-2.65v3.53h2.65a2.64 2.64 0 1 1 0 5.28H35v3.53Z"/><path fill="#fff" d="M38.54 21.89a6.17 6.17 0 0 1 6.17 6.17a6.1 6.1 0 0 1-.64 2.73a6.16 6.16 0 0 0-4.64-10.22h-2.65v1.32Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M39.43 32.91a6.17 6.17 0 1 0 0-12.34h-2.65v3.53h2.65a2.64 2.64 0 1 1 0 5.28H35v3.53Z"/><path fill="#fff" d="M38.54 16.6c0-1.7-6.51-3.08-14.54-3.08S9.46 14.9 9.46 16.6v8.82a14.54 14.54 0 0 0 29.08 0Z"/><path fill="#e8f4fa" d="M24 36.44a16.73 16.73 0 0 1-14.2-7.9a14.54 14.54 0 0 0 28.4 0a16.73 16.73 0 0 1-14.2 7.9Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M38.54 16.6c0-1.7-6.51-3.08-14.54-3.08S9.46 14.9 9.46 16.6v8.82a14.54 14.54 0 0 0 29.08 0Z"/><path fill="#bf8256" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.81 18.8A30 30 0 0 0 24 17a30 30 0 0 0-10.81 1.8A30.27 30.27 0 0 0 24 20.57a30.27 30.27 0 0 0 10.81-1.77Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M11.22 17.82c2.63 1.65 7.36 2.75 12.78 2.75s10.15-1.1 12.78-2.75"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.07 4.26a1.76 1.76 0 1 0 3.52 0a1.76 1.76 0 1 0-3.52 0Zm-4.28 2.5a4 4 0 1 0-3.23 1.67a4 4 0 0 0 .5 0a2.63 2.63 0 0 0-.06.49a2.21 2.21 0 1 0 2.79-2.11Z"/>`),
		g.Group(children),
	)
}