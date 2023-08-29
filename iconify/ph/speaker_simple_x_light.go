package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerSimpleXLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M154.64 26.61a6 6 0 0 0-6.32.65L77.94 82H32a14 14 0 0 0-14 14v64a14 14 0 0 0 14 14h45.94l70.38 54.74A6 6 0 0 0 158 224V32a6 6 0 0 0-3.36-5.39ZM146 211.73l-62.32-48.47A6 6 0 0 0 80 162H32a2 2 0 0 1-2-2V96a2 2 0 0 1 2-2h48a6 6 0 0 0 3.68-1.26L146 44.27Zm98.24-64a6 6 0 1 1-8.48 8.48L216 136.48l-19.76 19.76a6 6 0 0 1-8.48-8.48L207.52 128l-19.76-19.76a6 6 0 0 1 8.48-8.48L216 119.52l19.76-19.76a6 6 0 0 1 8.48 8.48L224.48 128Z"/>`),
		g.Group(children),
	)
}