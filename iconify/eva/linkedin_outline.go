package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedinOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaLinkedinOutline0"><g id="evaLinkedinOutline1"><path id="evaLinkedinOutline2" fill="currentColor" d="M20 22h-1.67a2 2 0 0 1-2-2v-5.37a.92.92 0 0 0-.69-.93a.84.84 0 0 0-.67.19a.85.85 0 0 0-.3.65V20a2 2 0 0 1-2 2H11a2 2 0 0 1-2-2v-5.46a6.5 6.5 0 1 1 13 0V20a2 2 0 0 1-2 2Zm-4.5-10.31a3.73 3.73 0 0 1 .47 0a2.91 2.91 0 0 1 2.36 2.9V20H20v-5.46a4.5 4.5 0 1 0-9 0V20h1.67v-5.46a2.85 2.85 0 0 1 2.83-2.85ZM6 22H4a2 2 0 0 1-2-2V10a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2ZM4 10v10h2V10Zm1-3a3 3 0 1 1 3-3a3 3 0 0 1-3 3Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/></g></g>`),
		g.Group(children),
	)
}