package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Findroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5c-4.89 0-20.66 28.58-18.25 33.4s34.13 4.77 36.51 0S28.9 5.5 24 5.5Zm12 29.21c-1.56 3.13-22.35 3.17-23.93 0S20.8 12.83 24 12.83s13.52 18.76 12 21.88Z"/><path fill="currentColor" d="M21.291 31.325a.75.75 0 1 1 .745-.755v.036a.744.744 0 0 1-.745.72Zm5.476 0a.75.75 0 1 1 .745-.755v.036a.744.744 0 0 1-.745.72Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.015 25.73h0a6.707 6.707 0 0 1 6.707 6.707h0a.954.954 0 0 1-.954.953H18.232a.954.954 0 0 1-.954-.954h0a6.707 6.707 0 0 1 6.707-6.706Zm-5.783-1l1.907 2.238m9.675-2.238l-1.899 2.238"/>`),
		g.Group(children),
	)
}