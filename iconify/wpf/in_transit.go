package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InTransit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M0 4v1h8v1H2v1H0v1h7v1H2v1H0v1h6v1H2v1H0v1h5v1H2v4c0 .6.4 1 1 1h1c.2-1.7 1.7-3.094 3.5-3.094S10.8 18.2 11 20h3c.6 0 1-.4 1-1V5c0-.6-.4-1-1-1H0zm17 4c-.6 0-1 .4-1 1v10c0 .5.4 1 1 1c.2-1.7 1.7-3.094 3.5-3.094S23.8 18.2 24 20h1c.6 0 1-.4 1-1v-5c0-.9-.813-2-.813-2L23 9c-.5-.6-.8-1-1.5-1H17zm2 2h2.406c.2 0 .407.188.407.188l2.093 2.906c.2.3.407.612.407.812H19V10zM7.5 18.188A2.321 2.321 0 0 0 5.187 20.5A2.321 2.321 0 0 0 7.5 22.813A2.321 2.321 0 0 0 9.813 20.5A2.321 2.321 0 0 0 7.5 18.187zm13 0a2.321 2.321 0 0 0-2.313 2.312a2.321 2.321 0 0 0 2.313 2.313a2.321 2.321 0 0 0 2.313-2.313a2.321 2.321 0 0 0-2.313-2.313z"/>`),
		g.Group(children),
	)
}