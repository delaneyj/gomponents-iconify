package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JeepThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 100h-12.77l-10-46.51a12.07 12.07 0 0 0-11.7-9.49H50.47a12.07 12.07 0 0 0-11.74 9.49L28.77 100H16a4 4 0 0 0 0 8h12v100a12 12 0 0 0 12 12h24a12 12 0 0 0 12-12v-28h104v28a12 12 0 0 0 12 12h24a12 12 0 0 0 12-12V108h12a4 4 0 0 0 0-8ZM46.56 55.16A4 4 0 0 1 50.47 52h155.06a4 4 0 0 1 3.91 3.16l9.61 44.84H37ZM68 208a4 4 0 0 1-4 4H40a4 4 0 0 1-4-4v-28h32Zm148 4h-24a4 4 0 0 1-4-4v-28h32v28a4 4 0 0 1-4 4Zm4-40h-72v-36a4 4 0 0 0-8 0v36h-24v-36a4 4 0 0 0-8 0v36H36v-64h184ZM60 140a8 8 0 1 1 8 8a8 8 0 0 1-8-8Zm120 0a8 8 0 1 1 8 8a8 8 0 0 1-8-8Z"/>`),
		g.Group(children),
	)
}