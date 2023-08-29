package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caravan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 112c0-44.2 35.8-80 80-80h336c88.4 0 160 71.6 160 160v160h32c17.7 0 32 14.3 32 32s-14.3 32-32 32H288c0 53-43 96-96 96s-96-43-96-96H80c-44.2 0-80-35.8-80-80V112zm320 240h128v-96h-32c-8.8 0-16-7.2-16-16s7.2-16 16-16h32v-64c0-17.7-14.3-32-32-32h-64c-17.7 0-32 14.3-32 32v192zM96 128c-17.7 0-32 14.3-32 32v64c0 17.7 14.3 32 32 32h128c17.7 0 32-14.3 32-32v-64c0-17.7-14.3-32-32-32H96zm96 336a48 48 0 1 0 0-96a48 48 0 1 0 0 96z"/>`),
		g.Group(children),
	)
}