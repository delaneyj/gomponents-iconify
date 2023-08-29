package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandEyeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M144 184a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm76-68v36a92 92 0 0 1-184 0V76a32 32 0 0 1 40-31v-1a32 32 0 0 1 60.79-14A32 32 0 0 1 180 60v25a32 32 0 0 1 40 31Zm-24 0a8 8 0 0 0-16 0a12 12 0 0 1-24 0V60a8 8 0 0 0-16 0v44a12 12 0 0 1-24 0V44a8 8 0 0 0-16 0v68a12 12 0 0 1-24 0V76a8 8 0 0 0-16 0v76a68 68 0 0 0 136 0Zm-68 12a55.8 55.8 0 0 0-49.89 30.54a12 12 0 0 0 21.37 10.92a32 32 0 0 1 57 0a12 12 0 0 0 10.73 6.54a11.83 11.83 0 0 0 5.45-1.32a12 12 0 0 0 5.23-16.14A55.8 55.8 0 0 0 128 128Z"/>`),
		g.Group(children),
	)
}