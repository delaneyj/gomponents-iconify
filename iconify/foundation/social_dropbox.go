package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M34.063 13.993L11.288 28.864l15.747 12.612L50 27.295zM11.288 54.087l22.775 14.869L50 55.653L27.035 41.476zM50 55.653l15.937 13.303l22.775-14.869l-15.747-12.611zm38.712-26.789L65.937 13.993L50 27.295l22.965 14.181z"/><path fill="currentColor" d="M50.046 58.517L34.063 71.778l-6.84-4.464v5.005l22.823 13.688l22.825-13.688v-5.005l-6.841 4.464z"/>`),
		g.Group(children),
	)
}