package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOffTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9.187H8.812c-.756 0-1.134 0-1.44.101a2 2 0 0 0-1.271 1.272C6 10.865 6 11.243 6 12c0 .756 0 1.134.1 1.44a2 2 0 0 0 1.273 1.271c.305.101.683.101 1.44.101h1.423c.28 0 .42 0 .55.036a.998.998 0 0 1 .32.153c.11.078.198.187.374.406l1.675 2.073c.874 1.08 1.31 1.62 1.693 1.68a1 1 0 0 0 .891-.316c.261-.286.261-.98.261-2.37V15m0-4.5V6.977c0-.936 0-1.404-.122-1.628a1 1 0 0 0-1.26-.445c-.235.097-.53.461-1.118 1.19l-.625.773M6 5l14 14"/>`),
		g.Group(children),
	)
}