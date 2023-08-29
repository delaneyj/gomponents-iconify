package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xododocs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.11 25.24V10s-12.83 1.65-12.83 6.33c0 5.06 18.22 7.75 18.22 7.75V40s-18.22-2.47-18.22-8.3c0-3.86 12.83-6.46 12.83-6.46ZM9.89 22.76V38s12.83-1.65 12.83-6.33c0-5.06-18.22-7.75-18.22-7.75V8.05s18.22 2.47 18.22 8.3c0 3.81-12.83 6.41-12.83 6.41Z"/>`),
		g.Group(children),
	)
}