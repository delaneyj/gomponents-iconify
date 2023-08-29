package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YiHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="22.115" height="32.823" x="12.988" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="11.057"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.898 43.5h22.204"/><circle cx="24" cy="16.214" r="5.611" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}