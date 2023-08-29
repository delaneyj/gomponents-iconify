package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileZipThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 148h-16a4 4 0 0 0-4 4v56a4 4 0 0 0 8 0v-12h12a24 24 0 0 0 0-48Zm0 40h-12v-32h12a16 16 0 0 1 0 32Zm-52-36v56a4 4 0 0 1-8 0v-56a4 4 0 0 1 8 0Zm-40.53 2l-28.58 50H88a4 4 0 0 1 0 8H56a4 4 0 0 1-3.47-6l28.58-50H56a4 4 0 0 1 0-8h32a4 4 0 0 1 3.47 6Zm119.36-68.83l-56-56A4 4 0 0 0 152 28H56a12 12 0 0 0-12 12v72a4 4 0 0 0 8 0V40a4 4 0 0 1 4-4h92v52a4 4 0 0 0 4 4h52v20a4 4 0 0 0 8 0V88a4 4 0 0 0-1.17-2.83ZM156 84V41.65L198.34 84Z"/>`),
		g.Group(children),
	)
}