package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 64C0 28.7 28.7 0 64 0h406.1c21.1 0 36.4 20.1 30.9 40.4L494.5 64H336c-8.8 0-16 7.2-16 16s7.2 16 16 16h149.8l-17.5 64H336c-8.8 0-16 7.2-16 16s7.2 16 16 16h123.6l-17.5 64H336c-8.8 0-16 7.2-16 16s7.2 16 16 16h97.5L416 352H160l-8.7-96H64c-35.3 0-64-28.7-64-64V64zm145.5 128L133.8 64H64v128h81.5zM144 384h288c26.5 0 48 21.5 48 48v32c0 26.5-21.5 48-48 48H144c-26.5 0-48-21.5-48-48v-32c0-26.5 21.5-48 48-48zm144 96a32 32 0 1 0 0-64a32 32 0 1 0 0 64z"/>`),
		g.Group(children),
	)
}