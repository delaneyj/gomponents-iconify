package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftTeamsLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M108 104a4 4 0 0 1-4 4H92v44a4 4 0 0 1-8 0v-44H72a4 4 0 0 1 0-8h32a4 4 0 0 1 4 4Zm120-11.26V152a36 36 0 0 1-35.44 36a60 60 0 0 1-113.13 0H40a12 12 0 0 1-12-12V80a12 12 0 0 1 12-12h62.07a36 36 0 0 1 66.48-27.36a28 28 0 0 1 35 43.36h15.69a8.75 8.75 0 0 1 8.76 8.74Zm-56.77-44.13A36 36 0 0 1 158.64 84H184a20 20 0 1 0-12.77-35.39ZM110.71 68H136a12 12 0 0 1 12 12v1.29A28 28 0 1 0 110.71 68ZM40 180h96a4 4 0 0 0 4-4V80a4 4 0 0 0-4-4H40a4 4 0 0 0-4 4v96a4 4 0 0 0 4 4Zm148-12V96a4 4 0 0 0-4-4h-36v84a12 12 0 0 1-12 12H88a52 52 0 0 0 100-20Zm32-75.26a.74.74 0 0 0-.74-.74h-24a11.8 11.8 0 0 1 .7 4v72a60.23 60.23 0 0 1-1.18 11.86A28 28 0 0 0 220 152Z"/>`),
		g.Group(children),
	)
}