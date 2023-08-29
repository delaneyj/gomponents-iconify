package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12 39h24"/><rect width="39" height="26" x="4.5" y="9" rx="2" ry="2"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.86 22l-9.72-5.596v11.192L29.86 22z"/>`),
		g.Group(children),
	)
}