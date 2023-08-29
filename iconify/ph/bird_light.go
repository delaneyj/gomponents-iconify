package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BirdLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M174 68a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm-50.16 63.39a6 6 0 0 0-8.45.77l-40 48a6 6 0 1 0 9.22 7.68l40-48a6 6 0 0 0-.77-8.45ZM238 80a6 6 0 0 1-2.67 5L214 99.21V120a102.12 102.12 0 0 1-102 102H8a6 6 0 0 1-4.69-9.75L98 93.9V76a58 58 0 0 1 113.51-16.87L235.33 75a6 6 0 0 1 2.67 5Zm-16.82 0L203 67.88a6 6 0 0 1-2.47-3.48A46 46 0 0 0 110 76v20a6 6 0 0 1-1.31 3.75L20.48 210H112a90.1 90.1 0 0 0 90-90V96a6 6 0 0 1 2.67-5Z"/>`),
		g.Group(children),
	)
}