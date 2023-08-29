package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toilet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M7.98 45.34a16.02 1.66 0 1 0 32.04 0a16.02 1.66 0 1 0-32.04 0Z" opacity=".15"/><path fill="#fff" d="M34.07 44.5H18.62a1 1 0 0 1-1-1.1V26.55h20.44a17 17 0 0 0-3 9.69v7.16a1.05 1.05 0 0 1-.99 1.1Z"/><path fill="#e0e0e0" d="M17.63 40.08a14.45 14.45 0 0 0 2.57.23a13.41 13.41 0 0 0 12-7.9a2.17 2.17 0 0 1 2-1.37h1.72a16.42 16.42 0 0 1 2.19-4.49H17.63Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.07 44.5H18.62a1 1 0 0 1-1-1.1V26.55h20.44a17 17 0 0 0-3 9.69v7.16a1.05 1.05 0 0 1-.99 1.1Z"/><path fill="#fff" d="M28.85 26.55v18H13.51a1.1 1.1 0 0 1-.75-1.91a13.84 13.84 0 0 0 3.88-5.68a35.14 35.14 0 0 0 1.65-10.36Z"/><path fill="#e0e0e0" d="M16.64 36.91a12.67 12.67 0 0 1-1.36 2.76a33.55 33.55 0 0 0 5.65.47a18.77 18.77 0 0 0 7.92-1.73V26.55H18.29a35.14 35.14 0 0 1-1.65 10.36Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.85 26.55v18H13.51a1.1 1.1 0 0 1-.75-1.91a13.84 13.84 0 0 0 3.88-5.68a35.14 35.14 0 0 0 1.65-10.36Z"/><path fill="#fff" d="M29.87 7.22h10.58v18.23a1.1 1.1 0 0 1-1.1 1.1H31a1.1 1.1 0 0 1-1.1-1.1V7.22h-.03Z"/><path fill="#e0e0e0" d="M29.87 7.22h10.59v3.24H29.87z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M29.87 7.22h10.58v18.23a1.1 1.1 0 0 1-1.1 1.1H31a1.1 1.1 0 0 1-1.1-1.1V7.22h-.03Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.95 4.46h12.43v2.76H28.95Z"/><path fill="#fff" d="M31.25 26.55c0 5.72-4.95 10.36-11 10.36c-8.4 0-13.37-4.64-13.37-10.36Z"/><path fill="#e0e0e0" d="M7.28 29.42h23.54a9.73 9.73 0 0 0 .43-2.87H6.83a9.42 9.42 0 0 0 .45 2.87Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M31.25 26.55c0 5.72-4.95 10.36-11 10.36c-8.4 0-13.37-4.64-13.37-10.36Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M7.21 23.79a1.39 1.39 0 1 0 0 2.76h22.66v-2.76ZM28.95 9.84h.92v3.87h0h-.92a1.1 1.1 0 0 1-1.1-1.1v-1.66a1.1 1.1 0 0 1 1.1-1.11Z"/>`),
		g.Group(children),
	)
}