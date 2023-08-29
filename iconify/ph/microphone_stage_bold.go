package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneStageBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 12a75.9 75.9 0 0 0-75.51 84.33l-68.58 93.52a19.89 19.89 0 0 0 2 26l14.29 14.29a19.89 19.89 0 0 0 26 2l93.52-68.58A76 76 0 1 0 168 12Zm52 76a51.66 51.66 0 0 1-7.75 27.27l-71.51-71.52A52 52 0 0 1 220 88ZM54.72 210.71l-9.43-9.43l56.19-76.63a76.46 76.46 0 0 0 29.87 29.87ZM116 88a51.63 51.63 0 0 1 7.75-27.27l71.51 71.51A52 52 0 0 1 116 88Z"/>`),
		g.Group(children),
	)
}