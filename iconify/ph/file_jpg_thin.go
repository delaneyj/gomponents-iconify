package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileJpgThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 148h-16a4 4 0 0 0-4 4v56a4 4 0 0 0 8 0v-12h12a24 24 0 0 0 0-48Zm0 40h-12v-32h12a16 16 0 0 1 0 32Zm92-4v16.87a4 4 0 0 1-1.11 2.77A26.11 26.11 0 0 1 192 212c-15.44 0-28-14.36-28-32s12.56-32 28-32a25.41 25.41 0 0 1 14.24 4.43a4 4 0 1 1-4.48 6.63A17.45 17.45 0 0 0 192 156c-11 0-20 10.77-20 24s9 24 20 24a17.87 17.87 0 0 0 12-4.82V188h-4a4 4 0 0 1 0-8h8a4 4 0 0 1 4 4ZM76 152v38a22 22 0 0 1-44 0a4 4 0 0 1 8 0a14 14 0 0 0 28 0v-38a4 4 0 0 1 8 0Zm134.83-66.83l-56-56A4 4 0 0 0 152 28H56a12 12 0 0 0-12 12v72a4 4 0 0 0 8 0V40a4 4 0 0 1 4-4h92v52a4 4 0 0 0 4 4h52v20a4 4 0 0 0 8 0V88a4 4 0 0 0-1.17-2.83ZM156 84V41.65L198.34 84Z"/>`),
		g.Group(children),
	)
}