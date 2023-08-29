package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindStream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m25 2l-1.414 1.414L26.172 6h-5.09a16.93 16.93 0 0 0-6.314 1.216l-4.279 1.711A14.93 14.93 0 0 1 4.92 10H2v2h2.918a16.93 16.93 0 0 0 6.314-1.216l4.279-1.711A14.93 14.93 0 0 1 21.08 8h5.09l-2.586 2.586L25 12l5-5Z"/><path fill="currentColor" d="m21 11l-1.414 1.414L22.172 15h-3.27a16.942 16.942 0 0 0-5.97 1.083l-2.566.962A14.946 14.946 0 0 1 5.099 18H2v2h3.099a16.942 16.942 0 0 0 5.969-1.083l2.566-.962A14.946 14.946 0 0 1 18.901 17h3.27l-2.585 2.586L21 21l5-5Z"/><path fill="currentColor" d="m17 20l-1.414 1.414L18.172 24h-1.575a16.988 16.988 0 0 0-5.377.872l-1.073.358a14.985 14.985 0 0 1-4.744.77H2v2h3.403a16.988 16.988 0 0 0 5.377-.872l1.073-.358a14.985 14.985 0 0 1 4.744-.77h1.575l-2.586 2.586L17 30l5-5Z"/>`),
		g.Group(children),
	)
}