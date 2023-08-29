package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ABetterCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="24" r="14.5"/><path d="M26.192 20.204h-4.384L19.616 24l2.192 3.796h4.384L28.384 24l-2.192-3.796zM28.384 24l5.901 10.221m-20.57-20.442L19.616 24m8.093-14.017l-5.901 10.221m4.384 7.592l-5.901 10.221m1.517-10.221H10.003m16.189-7.592h11.805"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}