package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungHealth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.202 29.664c-.877.5-4.34 3.213-4.34 3.213s5.341 4.674 5.884 5.425c0 0-1.544-7.512-1.544-8.638Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.375 19.982c-1.753 0-4.173 1.947-6.385 4.145s-7.804 7.039-7.804 9.167s11.059 10.183 14.898 10.183s9.014.375 12.019-1.878s8.22-7.47 8.43-8.388a12.9 12.9 0 0 0 .167-4.132a1.887 1.887 0 0 0-1.586-1.168c-.814 0-2.942.876-2.942 2.358a13.681 13.681 0 0 0 .542 2.817a60.388 60.388 0 0 1-5.55 3.86s.125-7.575-.27-9.452s-3.86-4.883-3.86-5.467s1.188-4.465 1.188-4.465s2.442.542 2.713-.94s.083-1.543.083-1.543a1.954 1.954 0 0 0 .564-1.586c-.167-.876-.668-.689-.647-1.19S32.873 4.5 26.029 4.5s-6.885 5.634-6.885 6.698s3.15 5.404 2.837 7.136s-.271 1.586-.271 1.586s-.692.062-1.335.062Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.978 13.764c-.772.23-5.153 6.594-7.95 6.594a5.598 5.598 0 0 1-3.672-1.94c2.462.146 5.592-4.632 7.992-4.987a12.461 12.461 0 0 1 3.63.333Z"/>`),
		g.Group(children),
	)
}