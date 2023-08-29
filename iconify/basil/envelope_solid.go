package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.29 4.908a54.397 54.397 0 0 1 9.42 0l1.511.13a2.889 2.889 0 0 1 2.313 1.546a.236.236 0 0 1-.091.307l-6.266 3.88a4.25 4.25 0 0 1-4.4.045L3.47 7.088a.236.236 0 0 1-.103-.293A2.889 2.889 0 0 1 5.78 5.039l1.51-.131Z"/><path fill="currentColor" d="M3.362 8.767a.248.248 0 0 0-.373.187a30.351 30.351 0 0 0 .184 7.56A2.888 2.888 0 0 0 5.78 18.96l1.51.131c3.135.273 6.287.273 9.422 0l1.51-.13a2.888 2.888 0 0 0 2.606-2.449a30.35 30.35 0 0 0 .161-7.779a.248.248 0 0 0-.377-.182l-5.645 3.494a5.75 5.75 0 0 1-5.951.061l-5.653-3.34Z"/>`),
		g.Group(children),
	)
}