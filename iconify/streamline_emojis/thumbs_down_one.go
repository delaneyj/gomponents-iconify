package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsDownOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="streamlineEmojisThumbsDown10" fill="#ffe500" d="M37.17 19.08h.18a2.55 2.55 0 0 1 0 5.09h-2.54a2.54 2.54 0 0 1 0 5.08h-6.12c.58 1.44 1.32 3.72 2 5.87a4.19 4.19 0 0 1-2.77 5.21a1.09 1.09 0 0 1-1.36-.73c-.86-2.93-1.5-6.65-6-10a18.61 18.61 0 0 1-4.05-4.35a3.91 3.91 0 0 0-3.19-1.69H9.79a1 1 0 0 1-1-1V11.44a1 1 0 0 1 1-1h2.49c3.28-.07 4.7-1.31 8.79-1.31h13.47a2.25 2.25 0 0 1 0 4.5h2.63a2.73 2.73 0 0 1 0 5.45Z"/></defs><path fill="#45413c" d="M13 45.5a11 1.5 0 1 0 22 0a11 1.5 0 1 0-22 0Z" opacity=".15"/><use href="#streamlineEmojisThumbsDown10"/><use href="#streamlineEmojisThumbsDown10"/><path fill="#fff48c" d="M9.79 13.71h2.49c3.28-.08 4.7-1.32 8.79-1.32h13.47a2.25 2.25 0 0 1 1.55.62a2.25 2.25 0 0 0-1.55-3.88H21.07c-4.09 0-5.51 1.24-8.79 1.31H9.79a1 1 0 0 0-1 1v3.27a1 1 0 0 1 1-1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m28.37 23.43l-.06-.11a11.39 11.39 0 0 0-8.69-6.22h0"/><path fill="#ffe500" d="M25.39 9.13h11.4v4.5h-11.4Z"/><path fill="#fff48c" d="M27.64 11.38h6.9a2.26 2.26 0 0 1 2 1.12a2.25 2.25 0 0 0-2-3.37h-6.9a2.25 2.25 0 0 0-1.95 3.37a2.26 2.26 0 0 1 1.95-1.12Z"/><path fill="#ffe500" d="M24.71 13.63h15.18v5.45H24.71Z"/><path fill="#ffe500" d="M24.71 19.08h15.18v5.09H24.71Zm.68 5.09h11.96v5.09H25.39Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M25.39 9.13h11.4v4.5h-11.4Zm-.68 4.5h15.18v5.45H24.71Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24.71 19.08h15.18v5.09H24.71Zm.68 5.09h11.96v5.09H25.39Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M37.17 19.08h.18a2.55 2.55 0 0 1 0 5.09h-2.54a2.54 2.54 0 0 1 0 5.08h-6.12c.58 1.44 1.32 3.72 2 5.87a4.19 4.19 0 0 1-2.77 5.21a1.09 1.09 0 0 1-1.36-.73c-.86-2.93-1.5-6.65-6-10a18.61 18.61 0 0 1-4.05-4.35a3.91 3.91 0 0 0-3.19-1.69H9.79a1 1 0 0 1-1-1V11.44a1 1 0 0 1 1-1h2.49c3.28-.07 4.7-1.31 8.79-1.31h13.47a2.25 2.25 0 0 1 0 4.5h2.63a2.73 2.73 0 0 1 0 5.45Z"/>`),
		g.Group(children),
	)
}