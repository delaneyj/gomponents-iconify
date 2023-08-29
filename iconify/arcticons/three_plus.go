package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.948 29.235v4.946m-2.473-2.472h4.946m-19.219-7.7h3.728M8.971 15.56c1.49-1.243 2.982-1.74 6.212-1.74h.746a4.985 4.985 0 0 1 4.97 4.97h0c0 2.734-2.237 5.165-4.97 5.165m-6.958 8.254c1.49 1.242 2.733 1.74 6.212 1.74h.746a4.985 4.985 0 0 0 4.97-4.97h0a4.985 4.985 0 0 0-4.97-4.97m19.881.497v6.439m-3.219-3.219h6.438m-8.07-13.244v9.323m-4.663-4.662h9.325"/>`),
		g.Group(children),
	)
}