package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Zm7.663-31.697C33.69 13.868 35 16.23 35 19.29c0 3.52-1.954 5.954-4.19 7.445c-1.527 1.02-3.258 1.66-4.81 1.994V30a2 2 0 1 1-4 0v-3a2 2 0 0 1 1.874-1.996c1.302-.082 3.2-.584 4.716-1.596C30.054 22.431 31 21.11 31 19.29c0-1.794-.714-2.997-1.782-3.822c-1.133-.875-2.776-1.396-4.614-1.461c-1.83-.065-3.673.33-5.103 1.08c-1.43.751-2.286 1.756-2.56 2.863a2 2 0 1 1-3.883-.96c.627-2.535 2.472-4.336 4.585-5.444c2.114-1.11 4.659-1.623 7.103-1.537c2.437.087 4.956.78 6.917 2.294ZM22 36a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}