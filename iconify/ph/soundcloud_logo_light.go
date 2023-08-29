package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundcloudLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M22 120v48a6 6 0 0 1-12 0v-48a6 6 0 0 1 12 0Zm26-30a6 6 0 0 0-6 6v96a6 6 0 0 0 12 0V96a6 6 0 0 0-6-6Zm32-8a6 6 0 0 0-6 6v104a6 6 0 0 0 12 0V88a6 6 0 0 0-6-6Zm32-32a6 6 0 0 0-6 6v136a6 6 0 0 0 12 0V56a6 6 0 0 0-6-6Zm109.06 57.88A78 78 0 0 0 144 42a6 6 0 0 0 0 12a65.75 65.75 0 0 1 65.67 59.33a6 6 0 0 0 4.83 5.29A34 34 0 0 1 208 186h-64a6 6 0 0 0 0 12h64a46 46 0 0 0 13.06-90.12Z"/>`),
		g.Group(children),
	)
}