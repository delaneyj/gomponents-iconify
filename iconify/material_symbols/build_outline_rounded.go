package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.875 20.375q-.5 0-.95-.188t-.825-.562l-5-5.025q-.5.2-1.012.3T9 15q-2.5 0-4.25-1.75T3 9q0-.9.25-1.712t.7-1.538L7.6 9.4l1.8-1.8l-3.65-3.65q.725-.45 1.538-.7T9 3q2.5 0 4.25 1.75T15 9q0 .575-.1 1.088t-.3 1.012l5.05 5q.375.375.563.825t.187.95q0 .5-.2.963t-.55.812q-.375.375-.825.55t-.95.175Zm-.325-2.125q.125.125.338.113t.337-.138q.125-.125.125-.338t-.125-.337L12.15 11.5q.45-.5.65-1.163T13 9q0-1.5-.963-2.613T9.65 5.05L11.5 6.9q.3.3.3.7t-.3.7l-3.2 3.2q-.3.3-.7.3t-.7-.3L5.05 9.65q.225 1.425 1.338 2.388T9 13q.65 0 1.3-.2t1.175-.625l6.075 6.075Zm-5.825-6.525Z"/>`),
		g.Group(children),
	)
}