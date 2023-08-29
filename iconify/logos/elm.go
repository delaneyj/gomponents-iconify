package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="#1293D8"><path d="m8.529 256l119.399-119.399L247.327 256zM0 8.673l119.399 119.399L0 247.471z"/><path fill-opacity=".75" d="M136.595 0H256v119.405z"/><path d="m136.456 128.072l55.436 55.436l55.435-55.436l-55.435-55.435z"/><path fill-opacity=".75" d="M8.529.144h110.87l52.024 52.024H60.553zm174.834 63.964l-55.435 55.436l-55.436-55.436zm72.493 183.363l-55.436-55.435l55.436-55.435z"/></g>`),
		g.Group(children),
	)
}