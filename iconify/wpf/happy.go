package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Happy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M13 .188C5.924.188.187 5.923.187 13S5.925 25.813 13 25.813c7.076 0 12.813-5.737 12.813-12.813C25.813 5.924 20.075.187 13 .187zm0 2c5.962 0 10.813 4.85 10.813 10.812S18.962 23.813 13 23.813C7.038 23.813 2.187 18.962 2.187 13C2.188 7.038 7.038 2.187 13 2.187zM9 9a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm8 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4zM5.75 16.094a8.83 8.83 0 0 0 7.25 3.75a8.829 8.829 0 0 0 7.25-3.75A12.374 12.374 0 0 1 13 18.437c-2.707 0-5.208-.874-7.25-2.343z"/>`),
		g.Group(children),
	)
}