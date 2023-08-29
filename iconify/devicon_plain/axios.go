package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Axios(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="m34 43.978l27.38-22.913l.038 91.495l-9.32 7.74l-.153-76.091zm62.962 38.345l-27.38 22.912l-.038-91.495L78.863 6l.154 76.091z"/>`),
		g.Group(children),
	)
}