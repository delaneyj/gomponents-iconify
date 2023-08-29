package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12h6M7.48 15.406l1.676 2.074c.873 1.08 1.31 1.62 1.692 1.68a1 1 0 0 0 .891-.316c.261-.286.261-.98.261-2.37V7.525c0-1.39 0-2.084-.26-2.37a1 1 0 0 0-.892-.315c-.383.059-.82.599-1.692 1.68L7.48 8.593c-.176.218-.264.327-.373.406a1 1 0 0 1-.32.153c-.13.035-.27.035-.551.035H4.813c-.757 0-1.135 0-1.44.101A2 2 0 0 0 2.1 10.56c-.1.305-.1.683-.1 1.44c0 .756 0 1.134.1 1.44a2 2 0 0 0 1.273 1.271c.305.101.683.101 1.44.101h1.423c.28 0 .42 0 .55.036a1 1 0 0 1 .32.152c.11.079.198.188.374.406Z"/>`),
		g.Group(children),
	)
}