package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdmBrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.388 11.012c5.289-.285 10.791 2.763 12.617 9.736a8.159 8.159 0 0 1 .06 16.258H13.26c-10.922-1.482-12.416-17.385 0-19.52a12.025 12.025 0 0 1 10.127-6.474Zm20.112 2.81h-5.656m2.828 2.827v-5.655"/><circle cx="24" cy="25.636" r="7.654" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.346 25.636h15.308M17.39 21.778h13.221m-13.22 7.715h13.22m-6.61 3.797V17.981"/><ellipse cx="24" cy="25.636" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.857" ry="7.654"/>`),
		g.Group(children),
	)
}