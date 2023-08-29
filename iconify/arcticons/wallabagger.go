package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallabagger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 38.62c1.29-.8 8.23-4.49 19.44.14c11.52 4.75 18.67.62 19.56.06a39.11 39.11 0 0 0-9.76-11.24l.75-.22c4.86-1.4 6.1-5.52 6.36-9a31.4 31.4 0 0 1 1.36-7.65c.83-3.45-.1-4.14-1-3.45c-.46.38-4.48 1.31-6.86 3.16c-3.85 3-6.22 8.71-7.29 11.86l-.22.64a1.75 1.75 0 0 1-1.5 1.21a11.64 11.64 0 0 0-1.4-.13c-.44 0-.87 0-1.31.06h0c-1.28.19-2-1.37-2.13-1.74C19 18.05 15 9.62 6 7.35c0 0-1.64-1.25-1.14.87s1.46 4.07 1.24 7c-.1 1.37-1 8.42 5.52 12.09c.62.35 1.16.64 1.66.88A40.69 40.69 0 0 0 4.5 38.61Z"/>`),
		g.Group(children),
	)
}