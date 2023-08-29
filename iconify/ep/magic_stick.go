package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagicStick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 64h64v192h-64V64zm0 576h64v192h-64V640zM160 480v-64h192v64H160zm576 0v-64h192v64H736zM249.856 199.04l45.248-45.184L430.848 289.6L385.6 334.848L249.856 199.104zM657.152 606.4l45.248-45.248l135.744 135.744l-45.248 45.248L657.152 606.4zM114.048 923.2L68.8 877.952l316.8-316.8l45.248 45.248l-316.8 316.8zM702.4 334.848L657.152 289.6l135.744-135.744l45.248 45.248L702.4 334.848z"/>`),
		g.Group(children),
	)
}