package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m349.1 140.6l-69.8 139.6l-46.5-46.5l-46.5 46.5l-46.5-46.5l-46.7 139.7h325.8l-69.8-232.8zM0 1v512h512V1H0zm465.5 418.9h-419V47.5h418.9v372.4zM139.6 187.2c25.7 0 46.5-20.9 46.5-46.5c0-25.7-20.9-46.5-46.5-46.5S93.1 115 93.1 140.6c0 25.7 20.8 46.6 46.5 46.6z"/>`),
		g.Group(children),
	)
}