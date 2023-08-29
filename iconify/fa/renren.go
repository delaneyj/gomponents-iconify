package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Renren(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1133 1442q-171 94-368 94q-196 0-367-94q138-87 235.5-211T765 963q35 144 132.5 268t235.5 211zM638 14v485q0 252-126.5 459.5T181 1265Q0 1050 0 770q0-187 83.5-349.5T313 151T638 14zm898 756q0 280-181 495q-204-99-330.5-306.5T898 499V14q179 30 325 137t229.5 269.5T1536 770z"/>`),
		g.Group(children),
	)
}