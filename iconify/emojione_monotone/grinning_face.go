package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrinningFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.43 2 2 15.43 2 32s13.43 30 30 30s30-13.43 30-30S48.57 2 32 2zm0 57.5C16.836 59.5 4.5 47.164 4.5 32S16.836 4.5 32 4.5S59.5 16.836 59.5 32S47.164 59.5 32 59.5z"/><path fill="currentColor" d="M47.783 34.006H16.215c-.785 0-1.505-.09-1.912.43C10.394 39.422 14.993 54 31.999 54c17.007 0 21.606-14.578 17.696-19.564c-.406-.52-1.127-.43-1.912-.43M32 52.182c-2.913 0-5.483-.561-7.713-1.517c2.152-1.032 4.711-1.663 7.713-1.663s5.563.631 7.713 1.663c-2.229.956-4.799 1.517-7.713 1.517M47.006 40H16.994c-2.001 0-2.001-4 .053-4h29.906c2.053 0 2.053 4 .053 4"/><circle cx="20.5" cy="23" r="5" fill="currentColor"/><circle cx="43.5" cy="23" r="5" fill="currentColor"/>`),
		g.Group(children),
	)
}