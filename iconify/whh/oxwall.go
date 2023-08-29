package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oxwall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm0-832q-87 0-160.5 43T235 351.5T192 512t43 160.5T351.5 789T512 832q85 0 159-42l145 26l-26-145q42-74 42-159q0-87-43-160.5T672.5 235T512 192z"/>`),
		g.Group(children),
	)
}