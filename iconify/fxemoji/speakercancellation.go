package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speakercancellation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#ADB8BC" d="M306.353 438.735L155.035 285.036c-18.786-19.082-18.786-49.708 0-68.79l151.318-153.7c30.772-31.257 83.962-9.467 83.962 34.395V404.34c0 43.862-53.19 65.652-83.962 34.395z"/><path fill="#597B91" d="M227.291 359.044h-89.869c-15.84 0-28.8-12.96-28.8-28.8V171.038c0-15.84 12.96-28.8 28.8-28.8h89.869v216.806z"/><path fill="#FF473E" d="M466.652 485.401a18.443 18.443 0 0 1-13.081-5.419L38.739 65.152c-7.225-7.225-7.225-18.938 0-26.163s18.938-7.225 26.163 0L479.733 453.82c7.226 7.225 7.226 18.938 0 26.162a18.437 18.437 0 0 1-13.081 5.419z"/>`),
		g.Group(children),
	)
}