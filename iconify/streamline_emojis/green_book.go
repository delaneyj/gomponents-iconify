package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreenBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#00f5bc" d="M38.15 40.91L15.78 34.2V1.92L38 8.58a2.4 2.4 0 0 1 1.7 2.29v28.9a1.19 1.19 0 0 1-1.55 1.14Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M38.15 40.91L15.78 34.2V1.92L38 8.58a2.4 2.4 0 0 1 1.7 2.29v28.9a1.19 1.19 0 0 1-1.55 1.14Z"/><path fill="#656769" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m20.66 3.39l-4.88-1.47V34.2l4.88 1.46V3.39zm-4.88-1.47L9.8 6.11v13l5.98-5.31V1.92z"/><path fill="#45413c" d="M10.25 45.21a13.75 1.79 0 1 0 27.5 0a13.75 1.79 0 1 0-27.5 0Z" opacity=".15"/><path fill="#fffce5" d="m32.16 43.94l4.54-3.46V11.19l-4.54 3.13v29.62z"/><path fill="#fffef2" d="M32.16 14.32v4.77l4.54-3.14v-4.76l-4.54 3.13z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m32.16 43.94l4.54-3.46V11.19l-4.54 3.13v29.62z"/><path fill="#fffef2" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.78 4.91L11.3 8.05l20.92 6.27l4.48-3.13l-20.92-6.28z"/><path fill="#00f5bc" d="m32.17 45.09l-21.51-6.45a1.2 1.2 0 0 1-.86-1.15V6.11L32 12.77a2.38 2.38 0 0 1 1.7 2.29V44a1.19 1.19 0 0 1-1.53 1.09Z"/><path fill="#8cffe4" d="M32 12.77L9.8 6.11v3.51L32 16.28a2.38 2.38 0 0 1 1.7 2.29v-3.51a2.38 2.38 0 0 0-1.7-2.29Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m32.17 45.09l-21.51-6.45a1.2 1.2 0 0 1-.86-1.15V6.11L32 12.77a2.38 2.38 0 0 1 1.7 2.29V44a1.19 1.19 0 0 1-1.53 1.09Z"/><path fill="#656769" d="M14.69 7.57L9.8 6.11v31.38a1.2 1.2 0 0 0 .86 1.15l4 1.21Z"/><path fill="#87898c" d="M9.8 6.11v3.51l4.89 1.46V7.57L9.8 6.11z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.69 7.57L9.8 6.11v31.38a1.2 1.2 0 0 0 .86 1.15l4 1.21Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m27.81 22.98l-7.62-2.24v-4.79l7.62 2.25v4.78z"/>`),
		g.Group(children),
	)
}