package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkyNews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.74 18.776v4.1a3.037 3.037 0 0 1-3.037 3.037h0c-.838 0-1.598-.34-2.147-.89"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.74 13.765v5.011a3.037 3.037 0 0 1-3.037 3.037h0a3.037 3.037 0 0 1-3.037-3.037v-5.01m-8.732-4.084V21.83m.001-2.581l5.499-5.472m-3.749 3.731l4.323 4.303M12.45 21.15c.555.466 1.154.68 2.499.68h.681a2.01 2.01 0 0 0 2.008-2.012h0a2.01 2.01 0 0 0-2.008-2.012h-1.363a2.01 2.01 0 0 1-2.007-2.012h0a2.01 2.01 0 0 1 2.007-2.012h.682c1.344 0 1.943.213 2.498.679M31.16 28.662l-1.761 5.656l-1.761-5.656l-1.761 5.656l-1.761-5.656m9.238 5.179c.39.327.811.477 1.756.477h.479c.78 0 1.411-.633 1.411-1.414h0c0-.78-.632-1.414-1.411-1.414h-.958c-.78 0-1.411-.633-1.411-1.414h0c0-.781.632-1.414 1.411-1.414h.479c.945 0 1.366.15 1.756.477M15.27 34.318v-3.522a2.135 2.135 0 0 0-2.135-2.134h0A2.135 2.135 0 0 0 11 30.796m0 3.522v-5.656m10.755 4.579a2.134 2.134 0 0 1-1.855 1.077h0a2.135 2.135 0 0 1-2.135-2.134v-1.388c0-1.179.956-2.134 2.135-2.134h0c1.179 0 2.134.955 2.134 2.134v.694h-4.269"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}