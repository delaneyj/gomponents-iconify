package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q.65 0 1.075-.425T13.5 16.5q0-.65-.425-1.075T12 15q-.65 0-1.075.425T10.5 16.5q0 .65.425 1.075T12 18Zm0-4.5q.65 0 1.075-.425T13.5 12q0-.65-.425-1.075T12 10.5q-.65 0-1.075.425T10.5 12q0 .65.425 1.075T12 13.5ZM12 9q.65 0 1.075-.425T13.5 7.5q0-.65-.425-1.075T12 6q-.65 0-1.075.425T10.5 7.5q0 .65.425 1.075T12 9Zm-5 6v-1.15q-1.275-.35-2.138-1.4T4 10h3V8.85q-1.275-.35-2.138-1.4T4 5h3q0-.825.588-1.413T9 3h6q.825 0 1.413.588T17 5h3q0 1.4-.863 2.45T17 8.85V10h3q0 1.4-.863 2.45T17 13.85V15h3q0 1.4-.863 2.45T17 18.85V19q0 .825-.588 1.413T15 21H9q-.825 0-1.413-.588T7 19v-.15q-1.275-.35-2.138-1.4T4 15h3Zm2 4h6V5H9v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}