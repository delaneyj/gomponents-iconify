package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagHaiti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#d22f27" d="M5 36h62v19H5z"/><path fill="#fff" d="M29 29h14v14H29z"/><path fill="none" stroke="#a57939" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M36 33v8"/><path fill="none" stroke="#a57939" stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="m40 37l-4 4m-4-4l4 4"/><path fill="none" stroke="#5c9e31" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M33 32h6"/><path fill="#5c9e31" d="M29 41h14v2H29z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}