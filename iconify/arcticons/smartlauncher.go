package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smartlauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.141 30.5c0 4.736 4.467 2.471 4.809 2.357c1.449-.487 9.263-5.438 10.545-6.089s5.47-2.332 5.104 2.947c-.321 4.634-5.135 6.872-5.104 6.898M25.95 22.172c-1.9 1.333-4.808 3.59-4.808 8.327M9.677 12.777c.85-.563 8.27-4.814 12.354-7.15a5.924 5.924 0 0 1 5.335.014l10.945 6.319c1.471.85 1.462 2.21-.018 3.045c-3.944 2.222-10.99 6.218-12.343 7.167"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.078 33.092s-11.699-6.995-12.15-7.28a8.454 8.454 0 0 1-3.535-6.53a6.594 6.594 0 0 1 3.285-6.505m22.532 5.78l4.35 2.511a9.742 9.742 0 0 1 5.042 7.497"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.131 28.898l-6.129 3.254a1.704 1.704 0 0 0-.053 2.98l12.392 7.154a5.329 5.329 0 0 0 5.328 0l9.827-5.673"/>`),
		g.Group(children),
	)
}