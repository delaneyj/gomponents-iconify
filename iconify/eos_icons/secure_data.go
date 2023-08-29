package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SecureData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.269 15.01a4.282 4.282 0 0 1 .78-2h-4.027v-2h7.994v.001a4.903 4.903 0 0 1 1.998.44V3.012a2.001 2.001 0 0 0-2.001-2.001H3.016a2.002 2.002 0 0 0-2.002 2.001V17.01a2.001 2.001 0 0 0 2.002 2.002h9.997v-.801a3.013 3.013 0 0 1 .268-1.2H7.013v-2Zm4.744-6h-7.994v-2h7.994Zm-12-6h12v2h-12Zm-2.011 14h-2v-2h2Zm0-12h-2v-2h2Zm2.005 2h2v2h-2Zm.003 6v-2h2v2Zm14.803 4.001v-1.5a2.818 2.818 0 0 0-5.6 0v1.5a1.29 1.29 0 0 0-1.2 1.2v3.5a1.31 1.31 0 0 0 1.2 1.3h5.5a1.31 1.31 0 0 0 1.3-1.2v-3.5a1.31 1.31 0 0 0-1.2-1.3Zm-1.3 0h-3v-1.5a1.375 1.375 0 0 1 1.5-1.3a1.375 1.375 0 0 1 1.5 1.3Z"/>`),
		g.Group(children),
	)
}