package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUpOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M13 45.5a11 1.5 0 1 0 22 0a11 1.5 0 1 0-22 0Z" opacity=".15"/><path fill="#ffe500" d="M37.17 30.42h.18a2.54 2.54 0 1 0 0-5.08h-2.54a2.55 2.55 0 0 0 0-5.09h-6.12c.58-1.44 1.32-3.71 2-5.87a4.18 4.18 0 0 0-2.77-5.2a1.08 1.08 0 0 0-1.36.72c-.86 2.94-1.5 6.65-6 10a18.23 18.23 0 0 0-4 4.34a4 4 0 0 1-3.2 1.7H9.79a1 1 0 0 0-1 1v11.12a1 1 0 0 0 1 1h2.49c3.28.08 4.7 1.31 8.79 1.31h13.47a2.25 2.25 0 0 0 0-4.5h2.63a2.73 2.73 0 0 0 0-5.45Z"/><path fill="#fff48c" d="M14.34 28.63a3.93 3.93 0 0 0 3.19-1.7a18 18 0 0 1 4-4.34c4.5-3.32 5.15-7 6-10a1.08 1.08 0 0 1 1.35-.72a4.09 4.09 0 0 1 1.91 1.23a4.19 4.19 0 0 0-2.95-4a1.08 1.08 0 0 0-1.36.72c-.86 2.94-1.5 6.65-6 10a18.23 18.23 0 0 0-4 4.34a4 4 0 0 1-3.2 1.7H9.79a1 1 0 0 0-1 1v1.72Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m28.37 26.07l-.06.12a11.41 11.41 0 0 1-8.68 6.22h0"/><path fill="#ffe500" d="M25.39 35.87h11.4v4.5h-11.4Z"/><path fill="#ffe500" d="M24.71 30.42h15.18v5.45H24.71Z"/><path fill="#ffe500" d="M24.71 25.34h15.18v5.09H24.71Zm.68-5.09h11.96v5.09H25.39Z"/><path fill="#fff48c" d="M27.93 22.34h6.88a2.56 2.56 0 0 1 2.32 1.5a2.6 2.6 0 0 0 .22-1.05a2.54 2.54 0 0 0-2.54-2.54h-6.88a2.54 2.54 0 0 0-2.54 2.54a2.6 2.6 0 0 0 .22 1.05a2.56 2.56 0 0 1 2.32-1.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M25.39 35.87h11.4v4.5h-11.4Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24.71 30.42h15.18v5.45H24.71Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24.71 25.34h15.18v5.09H24.71Zm.68-5.09h11.96v5.09H25.39Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M37.17 30.42h.18a2.54 2.54 0 1 0 0-5.08h-2.54a2.55 2.55 0 0 0 0-5.09h-6.12c.58-1.44 1.32-3.71 2-5.87a4.18 4.18 0 0 0-2.77-5.2a1.08 1.08 0 0 0-1.36.72c-.86 2.94-1.5 6.65-6 10a18.23 18.23 0 0 0-4 4.34a4 4 0 0 1-3.2 1.7H9.79a1 1 0 0 0-1 1v11.12a1 1 0 0 0 1 1h2.49c3.28.08 4.7 1.31 8.79 1.31h13.47a2.25 2.25 0 0 0 0-4.5h2.63a2.73 2.73 0 0 0 0-5.45Z"/>`),
		g.Group(children),
	)
}