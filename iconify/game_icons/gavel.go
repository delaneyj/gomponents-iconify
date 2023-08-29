package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gavel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M129.75 24.47L21.625 211.655l55.78 32.188L185.532 56.656L129.75 24.47zm55.97 69.25l-75.626 130.874L326.47 349.47l75.592-130.876L185.72 93.72zm83.468.686l-11.22 19.438l84.97 49.03l11.25-19.468l-85-49zM434.25 200.22L326.156 387.405l55.78 32.188l108.095-187.188l-55.78-32.187zm-270.53 76.905l-11.282 19.53l84.968 49l11.25-19.5l-84.937-49.03zm-3.095 45.844L61.312 494.81h67.157l82.28-142.968l-50.125-28.875z"/>`),
		g.Group(children),
	)
}