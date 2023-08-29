package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M8.71 45.01a15.29 1.99 0 1 0 30.58 0a15.29 1.99 0 1 0-30.58 0Z" opacity=".15"/><path fill="#fff48c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.34 6.25a2.66 2.66 0 1 0 5.32 0a2.66 2.66 0 1 0-5.32 0Z"/><path fill="#ffe500" d="M42.62 35.51c0-3.44-3.77-5.71-5.32-8.64s-.83-7.76-2.66-13.3S27.77 6.92 24 6.92S15.19 8 13.36 13.57s-1.11 10.36-2.66 13.3s-5.32 5.2-5.32 8.64Z"/><path fill="#fff48c" d="M24 11.57c4.1 0 9.55 1.08 11.66 6.38a26.44 26.44 0 0 0-1-4.38C32.81 8 27.77 6.92 24 6.92S15.19 8 13.36 13.57a26.57 26.57 0 0 0-1 4.37c2.09-5.29 7.54-6.37 11.64-6.37Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M42.62 35.51c0-3.44-3.77-5.71-5.32-8.64s-.83-7.76-2.66-13.3S27.77 6.92 24 6.92S15.19 8 13.36 13.57s-1.11 10.36-2.66 13.3s-5.32 5.2-5.32 8.64Z"/><path fill="#ffe500" d="M5.38 35.51a18.62 1.99 0 1 0 37.24 0a18.62 1.99 0 1 0-37.24 0Z"/><path fill="#ebcb00" d="M24 33.52c-10.28 0-18.62.89-18.62 2c0 .36.91.71 2.5 1a102.44 102.44 0 0 1 16.12-1a102.44 102.44 0 0 1 16.12 1c1.59-.29 2.5-.64 2.5-1c0-1.11-8.34-2-18.62-2Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M5.38 35.51a18.62 1.99 0 1 0 37.24 0a18.62 1.99 0 1 0-37.24 0Z"/><path fill="#ffe500" d="M24 33.52h-3.46a4 4 0 1 0 6.92 0c-1.13.01-2.27 0-3.46 0Z"/><path fill="#ebcb00" d="M24 36.06h3.93a3.68 3.68 0 0 0 .06-.58a4 4 0 0 0-.53-2h-6.92a4 4 0 0 0-.53 2a3.68 3.68 0 0 0 .06.58c1.28.02 2.58 0 3.93 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 33.52h-3.46a4 4 0 1 0 6.92 0c-1.13.01-2.27 0-3.46 0Z"/>`),
		g.Group(children),
	)
}