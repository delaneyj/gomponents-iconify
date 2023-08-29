package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.82 4.687a10 10 0 0 1 .05 14.579M16.093 7.612a6 6 0 0 1 .03 8.748m-8.642-.953l1.676 2.073c.873 1.08 1.31 1.621 1.692 1.68a1 1 0 0 0 .891-.315c.261-.286.261-.981.261-2.37v-8.95c0-1.389 0-2.083-.26-2.37a1 1 0 0 0-.892-.315c-.383.059-.82.6-1.692 1.68L7.48 8.593c-.176.219-.264.328-.373.406a1 1 0 0 1-.32.153c-.13.036-.27.036-.551.036H4.813c-.757 0-1.135 0-1.44.1A2 2 0 0 0 2.1 10.56c-.1.306-.1.684-.1 1.44c0 .756 0 1.134.1 1.44a2 2 0 0 0 1.273 1.272c.305.1.683.1 1.44.1h1.423c.28 0 .42 0 .55.036a1 1 0 0 1 .32.153c.11.078.198.188.374.406Z"/>`),
		g.Group(children),
	)
}