package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Observation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M36 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 2a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm-17.904-2.895a3.771 3.771 0 0 0-5.333-.002l-.704.704a3.953 3.953 0 0 0-3.892 1.021a4.034 4.034 0 0 0 0 5.676l.833.839V34H6v2h3.05a3.5 3.5 0 1 0 4.899 0h20.102a3.5 3.5 0 1 0 4.899 0H42v-2h-4v-3h.066C40.24 31 42 29.224 42 27.033c0-2.19-1.761-3.967-3.934-3.967H20.11a.307.307 0 0 1-.218-.09l-.366-.37l.369-.367a3.77 3.77 0 0 0 .001-5.333l-1.8-1.8Zm.02 6.083l.365-.364a1.77 1.77 0 0 0 0-2.504l-1.8-1.8a1.771 1.771 0 0 0-2.504-.002l-.35.35l4.29 4.32ZM36 31v3H11v-8.644l5.275 5.311c.212.213.498.333.797.333H36Zm1-21a1 1 0 1 0-2 0v2.303l1.168 1.752a1 1 0 0 0 1.664-1.11L37 11.697V10Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}