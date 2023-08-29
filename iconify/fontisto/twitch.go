package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.481 6.259v6.259H9.39V6.259zm5.74 0v6.259H15.13V6.259zm0 10.962l3.649-3.663V2.091H3.649v15.13h4.702v3.13l3.13-3.13zM22.961 0v14.61l-6.259 6.259H12l-3.13 3.13H5.74V20.87H0V4.168L1.572 0z"/>`),
		g.Group(children),
	)
}