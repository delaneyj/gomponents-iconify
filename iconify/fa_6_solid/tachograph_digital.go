package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TachographDigital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 64C28.7 64 0 92.7 0 128v256c0 35.3 28.7 64 64 64h512c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64H64zm32 64h224c17.7 0 32 14.3 32 32v64c0 17.7-14.3 32-32 32H96c-17.7 0-32-14.3-32-32v-64c0-17.7 14.3-32 32-32zM64 368c0-8.8 7.2-16 16-16h256c8.8 0 16 7.2 16 16s-7.2 16-16 16H80c-8.8 0-16-7.2-16-16zm320 0c0-8.8 7.2-16 16-16h160c8.8 0 16 7.2 16 16s-7.2 16-16 16H400c-8.8 0-16-7.2-16-16zM80 288a16 16 0 1 1 0 32a16 16 0 1 1 0-32zm48 16a16 16 0 1 1 32 0a16 16 0 1 1-32 0zm80-16a16 16 0 1 1 0 32a16 16 0 1 1 0-32zm48 16a16 16 0 1 1 32 0a16 16 0 1 1-32 0zm80-16a16 16 0 1 1 0 32a16 16 0 1 1 0-32z"/>`),
		g.Group(children),
	)
}