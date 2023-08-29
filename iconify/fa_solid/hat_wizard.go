package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HatWizard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M496 448H16c-8.84 0-16 7.16-16 16v32c0 8.84 7.16 16 16 16h480c8.84 0 16-7.16 16-16v-32c0-8.84-7.16-16-16-16zm-304-64l-64-32l64-32l32-64l32 64l64 32l-64 32l-16 32h208l-86.41-201.63a63.955 63.955 0 0 1-1.89-45.45L416 0L228.42 107.19a127.989 127.989 0 0 0-53.46 59.15L64 416h144l-16-32zm64-224l16-32l16 32l32 16l-32 16l-16 32l-16-32l-32-16l32-16z"/>`),
		g.Group(children),
	)
}