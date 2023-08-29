package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 4.017C13 1.802 11.252 0 9.104 0H7.927C5.779 0 4.031 1.802 4.031 4.017V7.33a5.456 5.456 0 0 0 4.448 8.615a5.458 5.458 0 0 0 5.458-5.46A5.38 5.38 0 0 0 13 7.464V4.017zM9.104 2C10.15 2 11 2.904 11 4.017v1.646a5.398 5.398 0 0 0-2.521-.636c-.919 0-1.782.229-2.542.63V4.016C5.938 2.904 6.881 2 7.927 2h1.177zm-.625 9.846c.363 0 .693-.147.937-.379l2.578-.753a3.528 3.528 0 0 1-3.515 3.312a3.539 3.539 0 1 1 0-7.076A3.532 3.532 0 0 1 12 10.299l-2.573-.781a1.352 1.352 0 0 0-.947-.388a1.358 1.358 0 0 0-.001 2.716z"/>`),
		g.Group(children),
	)
}