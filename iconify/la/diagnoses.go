package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diagnoses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 5c-2.2 0-4 1.8-4 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4zm0 8h-1.38a5 5 0 0 0-4.33 2.52c.71.08 1.33.46 1.73 1.01c.52-.94 1.51-1.53 2.6-1.53h2.76c.44 0 .85.09 1.23.26c.21-.69.71-1.24 1.35-1.53a4.9 4.9 0 0 0-2.58-.73H16zm0-6c1.12 0 2 .88 2 2s-.88 2-2 2s-2-.88-2-2s.88-2 2-2zm5 8a1 1 0 0 0 0 2a1 1 0 0 0 0-2zm-11 2a1 1 0 0 0 0 2a1 1 0 0 0 0-2zm12.73.81c-.45.43-1.06.69-1.73.69c-.05 0-.11 0-.16-.01l1.25 2.92l.19.47l.53.09l5 1l.38-1.94l-4.47-.91l-.99-2.31zM15 18a1 1 0 0 0 0 2a1 1 0 0 0 0-2zm-6.63 1.9l-.09.22l-4.47.91l.38 1.94l5-1l.53-.09l.19-.47l.4-.93a2.514 2.514 0 0 1-1.941-.58zM18 21a1 1 0 0 0 0 2a1 1 0 0 0 0-2zM2 25v2h28v-2H2z"/>`),
		g.Group(children),
	)
}