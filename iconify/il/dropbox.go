package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M395 133L161 286L0 150L235 4zm143 493q5 0 10-3l96-63v37L395 754L147 597v-37l96 63q4 3 10 3t12-4l130-109l131 109q6 4 12 4zm92-340L395 133L556 4l235 146zm0 0l143 115l-232 152l-146-121zM18 401l143-115l234 146l-145 121z"/>`),
		g.Group(children),
	)
}