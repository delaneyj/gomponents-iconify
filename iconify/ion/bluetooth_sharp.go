package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M397.41 161.13L236-.28v212.8l-94.17-80.72l-26 30.37L225.27 256L115.8 349.83l26 30.37l94.2-80.72v212.8l161.41-161.41L286.73 256ZM276 96.28l62.59 62.59L276 212.52Zm62.58 256.85L276 415.72V299.48Z"/>`),
		g.Group(children),
	)
}