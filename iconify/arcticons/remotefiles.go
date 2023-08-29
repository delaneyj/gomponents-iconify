package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remotefiles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.78 43.5l-5.25-4.56l5.25-4.64m3.58-12.07l5.25 4.56l-5.25 4.64m5.25-4.64H9.47"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.53 39h24a2 2 0 0 0 2-2V13.8l-9.28-9.28H14.58a2.05 2.05 0 0 0-2.05 2.05v12.55"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.92 4.5v9.3h9.61"/>`),
		g.Group(children),
	)
}