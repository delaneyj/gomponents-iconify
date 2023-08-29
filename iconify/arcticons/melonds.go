package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Melonds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43 27.67L27.67 43a5.18 5.18 0 0 1-7.34 0L5 27.67a5.18 5.18 0 0 1 0-7.34L20.33 5a5.18 5.18 0 0 1 7.34 0L43 20.33a5.18 5.18 0 0 1 0 7.34Z"/><rect width="21" height="10" x="9.21" y="14.71" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" transform="rotate(-45 19.707 19.717)"/><rect width="21" height="10" x="18.4" y="23.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" transform="rotate(-45 28.9 28.9)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.339 32.642l2.828-2.829l2.829 2.829l-2.829 2.828zm7.47-7.474l2.828-2.829l2.829 2.829l-2.829 2.828zm-14.701-3.69l6.364-6.364l2.828 2.828l-6.364 6.364z"/>`),
		g.Group(children),
	)
}