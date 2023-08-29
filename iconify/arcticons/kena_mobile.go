package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KenaMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.062 13.625a14.938 14.938 0 1 0 29.876 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.938 43.5a14.938 14.938 0 0 0-29.876 0"/><circle cx="24" cy="13.625" r="4.277" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.062 6.613h6.51m16.856 0h6.51L36.825 4.5"/>`),
		g.Group(children),
	)
}