package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Activitylauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="23.9" cy="38.16" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.8" ry="1.74"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.75 33h2.38c.61 1.76 1.37 2 2.86 1.61l1.17 2c-1.49.94-1.58 2.17 0 3.38l-1.09 1.81c-2.05-.77-2.79-.06-2.94 1.7h-2.38c-.16-2-1.28-2.3-3-1.7l-1.15-1.88c1.53-1.14 1.43-2.26 0-3.36l1.24-2.08c1.62.52 2.72.14 2.91-1.48Zm1.18-2.38l2-3.42a1.88 1.88 0 0 0-2-2.49a1.79 1.79 0 0 0-1.84 2.49ZM19.3 11.49a37.42 37.42 0 0 1 9.29 0A33.84 33.84 0 0 1 27.6 23a40.53 40.53 0 0 0-7.2 0a28.26 28.26 0 0 1-1.1-11.49Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.59 11.49a19 19 0 0 0-3.87-6.42c-.76-.71-.87-.81-1.59 0a18.43 18.43 0 0 0-3.83 6.42M27.6 23a14.92 14.92 0 0 1 2.86.77c-.06-1.38-.38-3-2.27-3.52m-8.48 0c-1.89.49-2.11 2.12-2.17 3.5A14.92 14.92 0 0 1 20.4 23"/>`),
		g.Group(children),
	)
}