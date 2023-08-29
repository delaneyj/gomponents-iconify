package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TropicalDrink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M12.84 45.33a11.16 1.67 0 1 0 22.32 0a11.16 1.67 0 1 0-22.32 0Z" opacity=".15"/><path fill="#00b8f0" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.07 43.36c-.55-1.06-3.33-1.91-6.94-2.16a.87.87 0 0 1-.82-.86v-7.07h-2.62v7.07a.87.87 0 0 1-.82.86c-3.61.25-6.39 1.1-6.94 2.16a.41.41 0 0 0 .2.56c.85.44 3.58 1.58 8.87 1.58s8-1.14 8.87-1.58a.41.41 0 0 0 .2-.56Z"/><path fill="#00b8f0" d="M31.86 9.67c0-1-3.52-1.74-7.86-1.74s-7.86.78-7.86 1.74l-1.32 16.17a9.18 9.18 0 0 0 18.36 0Z"/><path fill="#4acfff" d="m32.19 13.71l-.33-4c0-1-3.52-1.74-7.86-1.74s-7.86.78-7.86 1.74l-1.32 16.13a9.23 9.23 0 0 0 1 4.22a33.55 33.55 0 0 0 9.9-7.28a43.17 43.17 0 0 0 6.47-9.07Z"/><path fill="#00b8f0" d="M16.14 9.67a7.86 1.75 0 1 0 15.72 0a7.86 1.75 0 1 0-15.72 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M31.86 9.67c0-1-3.52-1.74-7.86-1.74s-7.86.78-7.86 1.74l-1.32 16.17a9.18 9.18 0 0 0 18.36 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M31.86 9.67c0 1-3.52 1.75-7.86 1.75s-7.86-.78-7.86-1.75"/><path fill="#6dd627" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M12.49 12.16a.87.87 0 0 1-1.27-.4a5.68 5.68 0 0 1 9.51-5.88a.86.86 0 0 1-.21 1.31Z"/><path fill="#ff8a14" d="M17.88 16.23L17 25.84a7 7 0 1 0 14 0l-.87-9.61Z"/><path fill="#ff6242" d="M24 25.84a21.15 21.15 0 0 1-6.91-.91l-.08.91a7 7 0 1 0 14 0l-.08-.91a21.15 21.15 0 0 1-6.93.91Z"/><path fill="#ff8a14" d="M17.88 16.23a6.12.87 0 1 0 12.24 0a6.12.87 0 1 0-12.24 0Z"/><path fill="#bf8df2" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M26.18 17.54h2.63v-6.48a25.76 25.76 0 0 1-2.63.29Z"/><path fill="#ffaa54" d="M30.12 16.23c0 .48-2.74.87-6.12.87s-6.12-.39-6.12-.87L17 25.84a6.92 6.92 0 0 0 .78 3.21a33.16 33.16 0 0 0 8-6.27a47.05 47.05 0 0 0 4.44-5.7Z"/><path fill="#ff866e" d="m17.09 24.93l-.08.91a6.92 6.92 0 0 0 .78 3.21a36.48 36.48 0 0 0 4.74-3.24a18.26 18.26 0 0 1-5.44-.88Z"/><path fill="#9f5ae5" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m29.82 6.28l3.36-4L31 .5l-3.88 4.93a4.34 4.34 0 0 0-.94 2.7v3.22a25.76 25.76 0 0 0 2.63-.29v-2a4.31 4.31 0 0 1 1.01-2.78Z"/>`),
		g.Group(children),
	)
}