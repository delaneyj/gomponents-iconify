package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiChannelRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.05 21q-.45 0-.725-.313t-.2-.737q.725-5.275 2.087-8.613T8 8q1.2 0 2.1 1.688t1.5 3.887q.65-5.125 1.775-7.85T16 3q1.775 0 3.175 4.763t1.775 12.162q.05.425-.262.75T19.9 21q-.375 0-.687-.25t-.413-.625q-.55-2.05-1.388-3.588T16 15q-.575 0-1.4 1.5t-1.375 3.525q-.125.4-.463.688T12 21q-.425 0-.75-.275t-.375-.7q-.5-3.65-1.338-6.412T8 10.125q-.725.75-1.55 3.513t-1.325 6.437q-.05.375-.363.65T4.05 21Zm9.45-6.5q.5-.65 1.125-1.075T16 13q.725 0 1.363.425T18.5 14.5q-.425-3.575-1.1-6.213T16 5.05q-.725.625-1.4 3.238T13.5 14.5Z"/>`),
		g.Group(children),
	)
}