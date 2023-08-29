package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sanitizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5a1 1 0 0 1-.707-1.707l.829-.829A4.967 4.967 0 0 1 8.657 1H17a1 1 0 0 1 0 2H8.657a3.022 3.022 0 0 0-2.121.878l-.829.829A.997.997 0 0 1 5 5Z"/><path fill="currentColor" d="M10 3v2.5l.4-.3A1 1 0 0 1 11 5h4a1 1 0 0 1 .6.2l.4.3V3Z" opacity=".5"/><circle cx="13" cy="15" r="2" fill="currentColor"/><path fill="currentColor" d="M13 18a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3Z"/><path fill="currentColor" d="m18.8 7.6l-3.2-2.4A1 1 0 0 0 15 5h-4a1 1 0 0 0-.6.2L7.2 7.6A3.016 3.016 0 0 0 6 10v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V10a3.015 3.015 0 0 0-1.2-2.4ZM13 18a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3Z" opacity=".25"/>`),
		g.Group(children),
	)
}