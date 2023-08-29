package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M477 30H36C15.498 30 0 45.498 0 66v378c0 20.503 15.497 38 36 38h439c20.503 0 37-17.497 37-38V66c0-20.503-14.497-36-35-36zM285.834 244.425l49.182 43.516v-132.5h40v132.501l49.182-43.516l26.506 29.957l-95.688 84.665l-95.688-84.666l26.506-29.957zM97.185 281.127l-28.284-28.285l70.834-70.832l-70.834-70.836L97.185 82.89l99.119 99.119l-99.119 99.117zM459 433H252v-40h207v40z"/>`),
		g.Group(children),
	)
}