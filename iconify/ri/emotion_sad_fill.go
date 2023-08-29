package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmotionSadFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10a9.959 9.959 0 0 1-1.065 4.496a1.975 1.975 0 0 0-.398-.775l-.123-.135L19 14.172l-1.414 1.414l-.117.127a2 2 0 0 0 1.679 3.282A9.974 9.974 0 0 1 12 22C6.477 22 2 17.523 2 12S6.477 2 12 2Zm0 13c-1.38 0-2.63.56-3.534 1.463l-.166.174l.945.86C10.035 17.182 10.982 17 12 17c.905 0 1.754.144 2.486.396l.269.1l.945-.86A4.987 4.987 0 0 0 12 15Zm-3.5-5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm7 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`),
		g.Group(children),
	)
}