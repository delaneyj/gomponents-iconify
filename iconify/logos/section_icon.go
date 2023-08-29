package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SectionIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#00A86B" d="m98.184 302.837l121.62 25.912V101.305l-121.62 25.812v175.72Zm157.827 73.216L69.03 325.517V104.36l186.703-50.17l.267 321.874l.011-.01Z"/><path fill="#00A86B" d="M179.319 0L0 65.041v300.15l179.319 65.052V0ZM31.954 347.564V82.69l111.557-34.431v333.725l-111.557-34.43v.01Z"/><path fill="#038754" d="m179.319 109.88l17.693-3.744V69.973l-17.693 4.742v35.164Zm0 245.437l17.693 4.776v-36.197l-17.693-3.777v35.198Z"/>`),
		g.Group(children),
	)
}