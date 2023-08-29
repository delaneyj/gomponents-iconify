package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaSkipForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 111.882v288.236A23.979 23.979 0 0 0 70.285 419.4L248 287.763v112.355a23.979 23.979 0 0 0 38.285 19.282l194.56-144.119a24 24 0 0 0 0-38.57L286.285 92.6A24 24 0 0 0 248 111.882v112.355L70.285 92.6A24 24 0 0 0 32 111.882Zm248 15.881L453.119 256L280 384.237Zm-216 0L237.119 256L64 384.237Z"/>`),
		g.Group(children),
	)
}