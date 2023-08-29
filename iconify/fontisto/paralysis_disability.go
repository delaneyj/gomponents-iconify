package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParalysisDisability(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.878 17.986c-.924 2.33-3.15 3.951-5.758 3.968h-.005a6.08 6.08 0 0 1-3.62-10.965l.016-.012l-.355-2.191c-2.499 1.418-4.158 4.06-4.158 7.09a8.124 8.124 0 0 0 15.377 3.663l.021-.046z"/><path fill="currentColor" d="M16.248 15.07a1.747 1.747 0 0 0-.234-.194l-.006-.004c-.219-.15-.49-.24-.782-.24H8.895l-.819-5.047h6.758a.994.994 0 1 0 0-1.988H7.752l-.111-.672a1.411 1.411 0 1 0-2.783.47l-.001-.008l1.44 8.825a1.524 1.524 0 0 0 1.446 1.237h6.905l5.409 5.41a1.406 1.406 0 1 0 1.988-1.988zM8.345 2.691a2.691 2.691 0 1 1-5.383 0a2.691 2.691 0 0 1 5.383 0z"/>`),
		g.Group(children),
	)
}