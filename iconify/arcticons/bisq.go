package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bisq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.382 2.5C13.265 2.5 4.06 10.89 4.06 24.408S16.703 45.5 25.85 45.5s18.497-7.43 18.077-11.209c-.415-3.736-3.802-2.228-5.135-1.005s-7.931 6.205-12.389 6.205s-15.14-6.38-15.14-13.721S13.95 9.907 16.374 9.907c3.277 0 5.987 5.135 9.68 5.135s6.314-3.103 6.314-5.703S29.6 2.5 24.382 2.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.308 24c1.268 0 3.234.48 3.234 1.53c0 1.47-2.513 3.714-4.5 3.714s-2.077-1.814-2.077-2.622S30.774 24 33.308 24Zm-14.172 0c-1.268 0-3.234.48-3.234 1.53c0 1.47 2.513 3.714 4.5 3.714s2.077-1.814 2.077-2.622S21.67 24 19.136 24Zm20.137-1.573c1.989 0 2.622-1.158 2.622-2.447s-2.971-10.007-4.37-10.007s-6.38 5.156-6.38 7.21s5.703 5.244 8.128 5.244ZM25.516 35.806l-1.736-2.303a.831.831 0 0 1 .664-1.331h3.614a.831.831 0 0 1 .663 1.331l-1.735 2.303a.92.92 0 0 1-1.47 0Z"/>`),
		g.Group(children),
	)
}