package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#bdbec0" d="m11.022 34.332l3.536 3.535l-2.475 2.475a1 1 0 0 1-1.414 0l-2.122-2.121a1 1 0 0 1 0-1.414l2.475-2.475Z"/><path fill="#e0e0e0" d="m12.71 36l-3.18 3.2l-1-1a1 1 0 0 1 0-1.41L11 34.33Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m11.022 34.332l3.536 3.535h0l-2.475 2.475a1 1 0 0 1-1.414 0l-2.122-2.121a1 1 0 0 1 0-1.414l2.475-2.475h0Z"/><path fill="#87898c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m8.87 38.54l-2.8 2.81a1 1 0 0 0 1.41 1.41l2.8-2.8Z"/><path fill="#45413c" d="M8 43.5a12 1.5 0 1 0 24 0a12 1.5 0 1 0-24 0Z" opacity=".15"/><path fill="#c0dceb" d="M14.61 38L11 34.31a1 1 0 0 1-.09-1.31l15.67-18.32a1 1 0 0 1 1.47-.06l6.25 6.25a1 1 0 0 1-.06 1.47L16 38a1 1 0 0 1-1.39 0Z"/><path fill="#daedf7" d="m30 16.54l-1.92-1.92a1 1 0 0 0-1.47.06L10.91 33a1 1 0 0 0 .09 1.31L12.69 36Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.61 38L11 34.31a1 1 0 0 1-.09-1.31l15.67-18.32a1 1 0 0 1 1.47-.06l6.25 6.25a1 1 0 0 1-.06 1.47L16 38a1 1 0 0 1-1.39 0Z"/><path fill="#87898c" d="M25.5 14.83a8.5 8.5 0 1 0 17 0a8.5 8.5 0 1 0-17 0Z"/><path fill="#bdbec0" d="M28 12.82a8.5 8.5 0 0 1 12 0a8.39 8.39 0 0 1 2.24 4a8.49 8.49 0 1 0-16.5 0a8.39 8.39 0 0 1 2.26-4Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M25.5 14.83a8.5 8.5 0 1 0 17 0a8.5 8.5 0 1 0-17 0Z"/><path fill="#daedf7" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M40.36 19.78L29.05 8.47a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.41L39 21.19a1 1 0 0 0 1.41-1.41Z"/><path fill="#8ca4b8" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m24.802 22.606l1.415 1.415l-3.89 3.889l-1.414-1.415z"/><path fill="#e8f4fa" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m24.81 22.607l1.415 1.415l-1.768 1.767l-1.414-1.414z"/>`),
		g.Group(children),
	)
}