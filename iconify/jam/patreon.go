package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Patreon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.808.01c-3.95 0-7.164 3.196-7.164 7.125c0 3.916 3.214 7.103 7.164 7.103c3.938 0 7.142-3.187 7.142-7.103c0-3.93-3.204-7.125-7.142-7.125M.05 18.99V.01h3.502v18.98z"/>`),
		g.Group(children),
	)
}