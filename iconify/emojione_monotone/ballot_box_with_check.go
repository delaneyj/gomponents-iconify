package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallotBoxWithCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M14.632 30.707H9.368l14.737 23.925L62 2h-5.264L24.105 37.884z"/><path fill="currentColor" d="M52.526 54.632c0 .581-.471 1.052-1.053 1.052H9.368a1.053 1.053 0 0 1-1.053-1.052V12.526c0-.582.472-1.053 1.053-1.053h28.794l5.744-6.316H4.105A2.108 2.108 0 0 0 2 7.263v52.631C2 61.059 2.943 62 4.105 62h52.631a2.103 2.103 0 0 0 2.105-2.105V18.996l-6.315 8.773v26.863"/>`),
		g.Group(children),
	)
}