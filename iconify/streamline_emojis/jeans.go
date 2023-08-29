package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jeans(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M6.37 45.04a17.63 1.96 0 1 0 35.26 0a17.63 1.96 0 1 0-35.26 0Z" opacity=".15"/><path fill="#00b8f0" d="M37.38 4.18H10.62a.66.66 0 0 0-.66.65V42h11.76V20.82a2.28 2.28 0 0 1 4.56 0V42H38V4.83a.66.66 0 0 0-.62-.65Z"/><path fill="#4acfff" d="M37.38 4.18H10.62a.66.66 0 0 0-.66.65v6.53a.66.66 0 0 1 .66-.65h26.76a.66.66 0 0 1 .66.65V4.83a.66.66 0 0 0-.66-.65Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M37.38 4.18H10.62a.66.66 0 0 0-.66.65V42h11.76V20.82a2.28 2.28 0 0 1 4.56 0V42H38V4.83a.66.66 0 0 0-.62-.65Z"/><path fill="#4acfff" d="M9.31 38.78h13.06V44H9.31Zm16.32 0h13.06V44H25.63Z"/><path fill="#80ddff" d="M21.06 38.78H10.62a1.31 1.31 0 0 0-1.31 1.3v2.12a1.31 1.31 0 0 1 1.31-1.3h10.44a1.31 1.31 0 0 1 1.31 1.3v-2.12a1.31 1.31 0 0 0-1.31-1.3Zm16.32 0H26.94a1.31 1.31 0 0 0-1.31 1.3v2.12a1.31 1.31 0 0 1 1.31-1.3h10.44a1.31 1.31 0 0 1 1.31 1.3v-2.12a1.31 1.31 0 0 0-1.31-1.3Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.31 38.78h13.06V44H9.31Zm16.32 0h13.06V44H25.63Z"/><path fill="#4acfff" d="M21.72 4.18v8.07A1.72 1.72 0 0 0 23.44 14h2.19V4.18Z"/><path fill="#80ddff" d="M21.72 4.18h3.92v4.9h-3.92z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.72 4.18v8.07A1.72 1.72 0 0 0 23.44 14h2.19V4.18Z"/><path fill="#009fd9" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.47 4.18h3.92a.65.65 0 0 1 .65.65v8.49h0h-2.18a2.4 2.4 0 0 1-2.4-2.4V4.18h.01Zm-21.33 9.13H9.97h0V4.82a.65.65 0 0 1 .65-.65h3.88v6.74a2.4 2.4 0 0 1-2.36 2.4Z"/><path fill="#009fd9" d="M33.47 4.18h3.92a.65.65 0 0 1 .65.65v8.49h-2.18a2.4 2.4 0 0 1-2.4-2.4V4.18h.01Zm-21.33 9.13H9.97V4.82a.65.65 0 0 1 .65-.65h3.88v6.74a2.4 2.4 0 0 1-2.36 2.4Z"/><path fill="#00b8f0" d="M37.38 4.18h-3.91v3.26h3.91a.66.66 0 0 1 .66.65V4.83a.66.66 0 0 0-.66-.65Zm-26.76 0a.66.66 0 0 0-.66.65v3.26a.66.66 0 0 1 .66-.65h3.91V4.18Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.47 4.18h3.92a.65.65 0 0 1 .65.65v8.49h0h-2.18a2.4 2.4 0 0 1-2.4-2.4V4.18h.01Zm-21.33 9.13H9.97h0V4.82a.65.65 0 0 1 .65-.65h3.88v6.74a2.4 2.4 0 0 1-2.36 2.4Zm2.39-5.87h7.19m3.91 0h7.84"/>`),
		g.Group(children),
	)
}